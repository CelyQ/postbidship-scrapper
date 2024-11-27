package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"math"
	"strings"
	"time"
)

type Parser struct {
	app App
	ctx context.Context
}

func NewParser(ctx context.Context, app App) *Parser {
	return &Parser{ctx: ctx, app: app}
}

func (p *Parser) InputAccountCredentials() {
	if err := chromedp.Run(p.ctx, chromedp.Tasks{
		chromedp.Navigate(p.app.url),
		chromedp.WaitVisible("#signin", chromedp.ByID),
		chromedp.SendKeys("input[name=\"emailAddress\"]", p.app.email, chromedp.ByQuery),
		chromedp.Click("button[type=\"submit\"]", chromedp.ByQuery),
		chromedp.WaitVisible("input[type=\"password\"]", chromedp.ByQuery),
		chromedp.SendKeys("input[type=\"password\"]", p.app.password, chromedp.ByQuery),
		chromedp.Click("button[type=\"submit\"]", chromedp.ByQuery),
		chromedp.WaitVisible("a[href=\"/my-potential-loads\"]", chromedp.ByQuery),
		chromedp.Click("a[href=\"/my-potential-loads\"]", chromedp.ByQuery),
	}); err != nil {
		log.Fatal(err)
	}
}

func (p *Parser) GetPages() int {
	var text string
	if err := chromedp.Run(p.ctx, chromedp.Tasks{
		chromedp.Text(".datatable-footer > div", &text, chromedp.ByQuery),
	}); err != nil {
		log.Fatal(err)
	}

	var from, to, total int

	if _, err := fmt.Sscanf(text, "Showing %d to %d of %d entries", &from, &to, &total); err != nil {
		log.Fatal(err)
	}

	elementsPerPage := to - from + 1
	totalPages := int(math.Ceil(float64(total) / float64(elementsPerPage)))

	return totalPages
}

func (p *Parser) GetTableHeaders() []string {
	var headers []*cdp.Node

	if err := chromedp.Run(p.ctx, chromedp.Tasks{
		chromedp.WaitVisible(".rt-table[role=\"grid\"]", chromedp.ByQuery),
		chromedp.Nodes(".rt-th[role=\"columnheader\"] .rt-resizable-header-content > *", &headers, chromedp.ByQueryAll),
	}); err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, header := range headers {
		result = append(result, header.AttributeValue("title"))
	}

	return result
}

func (p *Parser) GetTableRows() [][]string {
	var tableGroups []*cdp.Node

	if err := chromedp.Run(p.ctx, chromedp.Tasks{
		chromedp.WaitNotPresent(".datatable-loader", chromedp.ByQuery),
		chromedp.Nodes(".rt-tr-group", &tableGroups, chromedp.ByQueryAll),
	}); err != nil {
		log.Fatal(err)
	}

	filteredTableGroups := filterTableGroups(p.ctx, tableGroups)

	var result [][]string

	for _, group := range filteredTableGroups {
		var cells []*cdp.Node

		if err := chromedp.Run(p.ctx, chromedp.Tasks{
			chromedp.Nodes("div[role=\"row\"] div[role=\"gridcell\"] > div", &cells, chromedp.ByQueryAll, chromedp.FromNode(group)),
		}); err != nil {
			log.Fatal(err)
		}

		var values []string

		for _, cell := range cells {
			value := strings.TrimSpace(cell.AttributeValue("title"))
			values = append(values, value)
		}

		result = append(result, values)
	}

	return result
}

func (p *Parser) NextPage() {
	err := chromedp.Run(p.ctx, chromedp.Tasks{
		chromedp.WaitNotPresent(".datatable-loader", chromedp.ByQuery),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click("li.btn-next-page > a", chromedp.ByQuery),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func filterTableGroups(ctx context.Context, groups []*cdp.Node) []*cdp.Node {
	var result []*cdp.Node

OUTER:
	for _, group := range groups {
		var rows []*cdp.Node

		if err := chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Nodes("div[role=\"row\"]", &rows, chromedp.ByQuery, chromedp.FromNode(group)),
		}); err != nil {
			log.Fatal(err)
		}

		for _, row := range rows {
			if strings.Contains(row.AttributeValue("class"), "-padRow") {
				continue OUTER
			}
		}

		result = append(result, group)
	}

	return result
}
