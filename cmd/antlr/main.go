package main

import (
	"antrl/internal/model"
	"antrl/internal/placeholder"
	"antrl/internal/replacer"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type bailErrorListener struct{ *antlr.DefaultErrorListener }

func (l bailErrorListener) SyntaxError(r antlr.Recognizer, s interface{}, line, col int, msg string, e antlr.RecognitionException) {
	panic(fmt.Errorf("syntax error at %d:%d: %s", line, col, msg))
}

func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//input := "777{{block_A}}8{{C}}8{{block_B}}88{{C}}8777{{block_A}}8{{C}}8{{block_B}}8777{{block_A}}8{{C}}8{{block_B}}88{{C}}8777{{block_A}}8{{C}}8{{block_B}}88{{C}}88{{end_B}}8{{C}}8{{end_A}}ssss8{{C}}8{{end_B}}8{{C}}8{{end_A}}sss777{{block_A}}8{{C}}8{{block_B}}88{{C}}8777{{block_A}}8{{C}}8{{block_B}}88{{C}}88{{end_B}}8{{C}}8{{end_A}}ssss8{{C}}8{{end_B}}8{{C}}8{{end_A}}sssss8{{C}}88{{end_B}}8{{C}}8{{end_A}}ssss8{{C}}8{{end_B}}8{{C}}8{{end_A}}ssss"
	input := "777{{block_A}}8{{C thj='1'}}8{{block_B}}88{{C}}87778{{C}}8{{end_B}}8777{{end_A}}dddd"
	//input := "777{{C}}8{{C}}8{{C}}88{{C}}87778{{C}}8{{C}}8777{{C}}dddd{{block_A}}s{{block_A}}sss{{block_A}}ssss{{end_A}}s{{end_A}}sss{{end_A}}"
	//input := antlr.NewInputStream("11111{{A}}22222{{A}}33333{{A}}44444{{A}}55555{{A}}666666{{A}}777{{block_B}}88{{block_B}}88{{F}}88{{end_B}}88{{end_B}}")
	//input := antlr.NewInputStream("aaa{{A  ty='t' }}weg{{A  ty='t' }}weg{{A  ty='t' }}weg{{A  ty='t' }}weg{{hehe}}weg")
	//input := antlr.NewInputStream("text1{{S}}erg{{block_B}}wegd{{A}}ws{{block_B}} wegd {{A}} weg{{end_B}} egrrrr{{end_B}}ss ")
	//input := antlr.NewInputStream("{{body attr='value' }}")
	//input := "{{block_A}}<section id=\"{{placeholder content='e3sgZ2xvYmFsYmxvY2tfYW5jaG9yIH19' uid='m2JtFEpE9h2o'}}_m2JtFEpE9h2o\" class=\"brz-section brz-css-i5k7np brz-css-jtcp0i\"><div class=\"brz-section__content brz-section--boxed brz-css-1tg3wgl brz-css-slitec\" data-brz-custom-id=\"nRy68eDiRIHR\"><div class=\"brz-container brz-css-1cx8gdk brz-css-thbmpo\"><div id=\"\" class=\"brz-css-rcv95 brz-css-1q9rhs4 brz-wrapper\"><div class=\"brz-wp-title brz-css-1f4xfm3 brz-css-1i0u9ei\" data-brz-custom-id=\"nnI1DEL6EOhD\"><span class=\"brz-wp-title-content\" style=\"min-height:20px\">{{placeholder content='e3ticml6eV9kY19wb3N0X3RpdGxlfX0='}}</span></div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-1muppmk brz-wrapper\"><div class=\"brz-menu__container brz-css-1yq8ne2 brz-css-1saf4ki\" data-mmenu-id=\"#tsdqcqfGPuEI_{{placeholder content='e3sgcmFuZG9tX2lkIH19' key='menu'}}\" data-mmenu-position=\"position-left\" data-mmenu-title=\"Menu\" data-mmenu-stickytitle=\"off\" data-mmenu-isslider=\"false\" data-mmenu-closingicon=\"%7B%22desktop%22%3A%22off%22%2C%22tablet%22%3A%22off%22%2C%22mobile%22%3A%22off%22%7D\" data-brz-custom-id=\"tsdqcqfGPuEI\">{{placeholder content='e3sgZ3JvdXAgfX0='}}<nav data-mods=\"%7B%22desktop%22%3A%22horizontal%22%2C%22tablet%22%3A%22horizontal%22%2C%22mobile%22%3A%22horizontal%22%7D\" class=\"brz-menu brz-menu__preview brz-css-19pdevo brz-css-1blhi8g\"><ul class=\"brz-menu__ul\">{{placeholder content='e3sgbWVudV9sb29wIH19' menuId='dc8d678facf2eea55abec859b6f8c26e3216a0b9a216aa60a6a8b5756e9ea4b0'}}<li data-menu-item-id=\"{{placeholder content='e3sgbWVudV9pdGVtX3VpZCB9fQ=='}}\" class=\"brz-menu__item {{placeholder content='e3sgbWVudV9pdGVtX2NsYXNzbmFtZSB9fQ=='}}\"><a class=\"brz-a\" target=\"{{placeholder content='e3sgbWVudV9pdGVtX3RhcmdldCB9fQ=='}}\" href=\"{{placeholder content='e3sgbWVudV9pdGVtX2hyZWYgfX0='}}\" title=\"{{placeholder content='e3sgbWVudV9pdGVtX2F0dHJfdGl0bGUgfX0='}}\">{{placeholder content='e3sgbWVudV9pdGVtX2ljb24gfX0='}}<span class=\"brz-span\">{{placeholder content='e3sgbWVudV9pdGVtX3RpdGxlIH19'}}</span></a>{{placeholder content='e3sgbWVnYV9tZW51IH19'}}{{placeholder content='e3sgbWVudV9sb29wX3N1Ym1lbnUgfX0='}}<ul class=\"brz-menu__sub-menu\">{{placeholder content='e3sgbWVudV9sb29wIH19' recursive='1'}}{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul>{{placeholder content='e3sgZW5kX21lbnVfbG9vcF9zdWJtZW51IH19'}}</li>{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul></nav>{{placeholder content='e3sgZW5kX2dyb3VwIH19'}}<div class=\"brz-mm-menu__icon\"><svg class=\"brz-icon-svg align-[initial]\"><use href=\"{{brizy_dc_url_site}}/icon/22459456/editor/menu-3.svg#brz_icon\"></use></svg></div>{{placeholder content='e3sgZ3JvdXAgfX0='}}<nav data-mods=\"%7B%22desktop%22%3A%22horizontal%22%2C%22tablet%22%3A%22horizontal%22%2C%22mobile%22%3A%22horizontal%22%7D\" id=\"tsdqcqfGPuEI_{{placeholder content='e3sgcmFuZG9tX2lkIH19' key='menu'}}\" class=\"brz-menu brz-menu__preview brz-menu__mmenu brz-menu--has-dropdown brz-css-n91ed brz-css-1akgap1\"><ul class=\"brz-menu__ul\">{{placeholder content='e3sgbWVudV9sb29wIH19' menuId='dc8d678facf2eea55abec859b6f8c26e3216a0b9a216aa60a6a8b5756e9ea4b0'}}<li data-menu-item-id=\"{{placeholder content='e3sgbWVudV9pdGVtX3VpZCB9fQ=='}}\" class=\"brz-menu__item {{placeholder content='e3sgbWVudV9pdGVtX2NsYXNzbmFtZSB9fQ=='}}\"><a class=\"brz-a\" target=\"{{placeholder content='e3sgbWVudV9pdGVtX3RhcmdldCB9fQ=='}}\" href=\"{{placeholder content='e3sgbWVudV9pdGVtX2hyZWYgfX0='}}\" title=\"{{placeholder content='e3sgbWVudV9pdGVtX2F0dHJfdGl0bGUgfX0='}}\">{{placeholder content='e3sgbWVudV9pdGVtX2ljb24gfX0='}}<span class=\"brz-span\">{{placeholder content='e3sgbWVudV9pdGVtX3RpdGxlIH19'}}</span></a>{{placeholder content='e3sgbWVnYV9tZW51IH19'}}{{placeholder content='e3sgbWVudV9sb29wX3N1Ym1lbnUgfX0='}}<ul class=\"brz-menu__sub-menu\">{{placeholder content='e3sgbWVudV9sb29wIH19' recursive='1'}}{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul>{{placeholder content='e3sgZW5kX21lbnVfbG9vcF9zdWJtZW51IH19'}}</li>{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul></nav>{{placeholder content='e3sgZW5kX2dyb3VwIH19'}}</div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-16qa2c6 brz-css-1rcpqcz brz-wrapper\"><div class=\"brz-posts brz-css-cl6qmp brz-css-11ckoy3\" data-brz-custom-id=\"kyxD0eOo0haq\"><div class=\"brz-posts__wrapper\">{{placeholder content='e3ticml6eV9kY19wb3N0X2xvb3B9fQ==' type='posts' collection_type='/collection_types/4959140' count='3' order_by='id' order='DESC' offset='0'}}<div class=\"brz-posts__item\"><div class=\"brz-columns brz-css-9izhcb brz-css-j2nle7 brz-css-tx3g8z\" data-brz-custom-id=\"klpJKINEfaiO\"><div class=\"brz-column__items brz-css-qbnzgb brz-css-rhzuxa\"><div id=\"\" class=\"brz-css-rcv95 brz-css-skwex9 brz-wrapper\"><div class=\"brz-image brz-css-15xpw96 brz-css-17lbhdj\" data-brz-custom-id=\"qmF9ME46aQeI\"><picture class=\"brz-picture brz-d-block brz-p-relative brz-css-gxdnbn brz-css-1ttms7k\"><source srcset=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='263' cW='350'}} 1x, {{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='526' cW='700'}} 2x\" media=\"(min-width: 992px)\"><source srcset=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='266' cW='354'}} 1x, {{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='532' cW='708'}} 2x\" media=\"(min-width: 768px)\"><img class=\"brz-img\" srcset=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='323' cW='430'}} 1x, {{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='646' cW='860'}} 2x\" src=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='263' cW='350'}}\" alt=\"{{placeholder content='e3sgYnJpenlfZGNfaW1hZ2VfYWx0IH19' entityId='' entityType='' imagePlaceholder='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ=='}}\" title=\"{{placeholder content='e3sgYnJpenlfZGNfaW1hZ2VfdGl0bGUgfX0=' entityId='' entityType='' imagePlaceholder='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ=='}}\" draggable=\"false\" loading=\"lazy\"></picture></div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-i7zumc brz-wrapper\"><div class=\"brz-wp-title brz-css-1f4xfm3 brz-css-1slhfhe brz-css-1jcy4c3\" data-brz-custom-id=\"yflMCbIEgZj7\"><span class=\"brz-wp-title-content\" style=\"min-height:20px\">{{placeholder content='e3ticml6eV9kY19wb3N0X3RpdGxlfX0='}}</span></div></div><div class=\"brz-wrapper-clone brz-flex-xs-wrap brz-css-kwv3lv brz-css-wckscm\" data-brz-custom-id=\"kTYuI3BOrgPm\"><a class=\"brz-a brz-btn brz-css-uk96b3 brz-css-z6u47w brz-css-tdg8w brz-css-yrfdqg brz-css-e6b431\" href=\"{{placeholder content='e3ticml6eV9kY191cmxfcG9zdH19'}}\" target=\"_self\" rel=\"noopener\" data-brz-link-type=\"external\" data-brz-custom-id=\"hPQBrYiaVrgJ\"><svg class=\"brz-icon-svg align-[initial] brz-css-w3n9bt brz-css-jxjvb7\"><use href=\"{{brizy_dc_url_site}}/icon/22459456/glyph/tail-right.svg#nc_icon\"></use></svg><span data-brz-translate-text=\"1\" class=\"brz-span brz-text__editor\">READ MORE</span></a></div></div></div></div>{{placeholder content='e3tlbmRfYnJpenlfZGNfcG9zdF9sb29wfX0='}}</div></div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-1oph5n4 brz-wrapper\"><div class=\"brz-post-info brz-css-1k7ppv3 brz-css-1endcla\" data-brz-custom-id=\"l9Pef9gQL8QE\"><div class=\"brz-ui-ed-list brz-ui-ed-list-split brz-ui-ed-post-info\"><div class=\"brz-ui-ed-spin-nested-loading\"><div class=\"brz-ui-ed-spin-container\"><ul class=\"brz-ui-ed-list-items\"><li class=\"brz-ui-ed-list-item brz-ui-ed-list-item-no-flex\"><svg class=\"brz-ui-ed-icon-svg\" width=\"1em\" height=\"1em\" viewBox=\"0 0 16 16\"><path fill=\"currentColor\" d=\"m13.985 11.439-3.063-.77c-.505-.126-.887-.491-.989-.944l-.187-.831c1.493-.6 2.452-1.91 2.454-3.354V3.85C12.235 1.815 10.44.118 8.129.002c-1.136-.03-2.238.345-3.054 1.04C4.26 1.738 3.8 2.695 3.8 3.694V5.54c.001 1.444.96 2.755 2.454 3.355l-.187.83c-.102.452-.484.817-.989.944l-3.063.769C1.414 11.589 1 12.072 1 12.622v2.764c0 .34.313.615.7.615h12.6c.387 0 .7-.275.7-.615v-2.764c0-.55-.414-1.032-1.015-1.183Z\"></path></svg>{{placeholder content='e3tjb2xsZWN0aW9uX2l0ZW1fYXV0aG9yfX0='}}</li><li class=\"brz-ui-ed-list-item brz-ui-ed-list-item-no-flex\"><svg class=\"brz-ui-ed-icon-svg\" width=\"1em\" height=\"1em\" viewBox=\"0 0 16 16\" fill=\"currentColor\"><g stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\"><path d=\"M16,3.5 L16,14.5 C16,15.3284505 15.2325149,16 14.2857143,16 L1.71428571,16 C0.767485119,16 0,15.3284505 0,14.5 L0,3.5 C0,2.67154948 0.767485119,2 1.71428571,2 L3.42857143,2 L3.42857143,0.375 C3.42857143,0.16796875 3.62053571,0 3.85714286,0 L5.28571429,0 C5.52232143,0 5.71428571,0.16796875 5.71428571,0.375 L5.71428571,2 L10.2857143,2 L10.2857143,0.375 C10.2857143,0.16796875 10.4776786,0 10.7142857,0 L12.1428571,0 C12.3794643,0 12.5714286,0.16796875 12.5714286,0.375 L12.5714286,2 L14.2857143,2 C15.2325149,2 16,2.67154948 16,3.5 Z M14,13.8223684 L14,5 L2,5 L2,13.8223684 C2,13.9204359 2.09161932,14 2.20454545,14 L13.7954545,14 C13.9083807,14 14,13.9204359 14,13.8223684 Z\" fill=\"currentColor\" fill-rule=\"nonzero\"></path></g></svg>{{placeholder content='e3tjb2xsZWN0aW9uX2l0ZW1fZGF0ZX19'}}</li><li class=\"brz-ui-ed-list-item brz-ui-ed-list-item-no-flex\"><svg viewBox=\"0 0 16 16\" width=\"1em\" height=\"1em\"><g fill=\"currentColor\"><path fill=\"currentColor\" d=\"M8,0C3.6,0,0,3.6,0,8s3.6,8,8,8s8-3.6,8-8S12.4,0,8,0z M12,9H7V4h2v3h3V9z\"></path></g></svg>{{placeholder content='e3tjb2xsZWN0aW9uX2l0ZW1fdGltZX19'}}</li></ul></div></div></div></div></div></div></div></section>{{end_A}}"
	//input := "{{block_A}}<section id=\"{{placeholder content='e3sgZ2xvYmFsYmxvY2tfYW5jaG9yIH19' uid='m2JtFEpE9h2o'}}_m2JtFEpE9h2o\" class=\"brz-section brz-css-i5k7np brz-css-jtcp0i\"><div class=\"brz-section__content brz-section--boxed brz-css-1tg3wgl brz-css-slitec\" data-brz-custom-id=\"nRy68eDiRIHR\"><div class=\"brz-container brz-css-1cx8gdk brz-css-thbmpo\"><div id=\"\" class=\"brz-css-rcv95 brz-css-1q9rhs4 brz-wrapper\"><div class=\"brz-wp-title brz-css-1f4xfm3 brz-css-1i0u9ei\" data-brz-custom-id=\"nnI1DEL6EOhD\"><span class=\"brz-wp-title-content\" style=\"min-height:20px\">{{placeholder content='e3ticml6eV9kY19wb3N0X3RpdGxlfX0='}}</span></div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-1muppmk brz-wrapper\"><div class=\"brz-menu__container brz-css-1yq8ne2 brz-css-1saf4ki\" data-mmenu-id=\"#tsdqcqfGPuEI_{{placeholder content='e3sgcmFuZG9tX2lkIH19' key='menu'}}\" data-mmenu-position=\"position-left\" data-mmenu-title=\"Menu\" data-mmenu-stickytitle=\"off\" data-mmenu-isslider=\"false\" data-mmenu-closingicon=\"%7B%22desktop%22%3A%22off%22%2C%22tablet%22%3A%22off%22%2C%22mobile%22%3A%22off%22%7D\" data-brz-custom-id=\"tsdqcqfGPuEI\">{{placeholder content='e3sgZ3JvdXAgfX0='}}<nav data-mods=\"%7B%22desktop%22%3A%22horizontal%22%2C%22tablet%22%3A%22horizontal%22%2C%22mobile%22%3A%22horizontal%22%7D\" class=\"brz-menu brz-menu__preview brz-css-19pdevo brz-css-1blhi8g\"><ul class=\"brz-menu__ul\">{{placeholder content='e3sgbWVudV9sb29wIH19' menuId='dc8d678facf2eea55abec859b6f8c26e3216a0b9a216aa60a6a8b5756e9ea4b0'}}<li data-menu-item-id=\"{{placeholder content='e3sgbWVudV9pdGVtX3VpZCB9fQ=='}}\" class=\"brz-menu__item {{placeholder content='e3sgbWVudV9pdGVtX2NsYXNzbmFtZSB9fQ=='}}\"><a class=\"brz-a\" target=\"{{placeholder content='e3sgbWVudV9pdGVtX3RhcmdldCB9fQ=='}}\" href=\"{{placeholder content='e3sgbWVudV9pdGVtX2hyZWYgfX0='}}\" title=\"{{placeholder content='e3sgbWVudV9pdGVtX2F0dHJfdGl0bGUgfX0='}}\">{{placeholder content='e3sgbWVudV9pdGVtX2ljb24gfX0='}}<span class=\"brz-span\">{{placeholder content='e3sgbWVudV9pdGVtX3RpdGxlIH19'}}</span></a>{{placeholder content='e3sgbWVnYV9tZW51IH19'}}{{placeholder content='e3sgbWVudV9sb29wX3N1Ym1lbnUgfX0='}}<ul class=\"brz-menu__sub-menu\">{{placeholder content='e3sgbWVudV9sb29wIH19' recursive='1'}}{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul>{{placeholder content='e3sgZW5kX21lbnVfbG9vcF9zdWJtZW51IH19'}}</li>{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul></nav>{{placeholder content='e3sgZW5kX2dyb3VwIH19'}}<div class=\"brz-mm-menu__icon\"><svg class=\"brz-icon-svg align-[initial]\"><use href=\"{{brizy_dc_url_site}}/icon/22459456/editor/menu-3.svg#brz_icon\"></use></svg></div>{{placeholder content='e3sgZ3JvdXAgfX0='}}<nav data-mods=\"%7B%22desktop%22%3A%22horizontal%22%2C%22tablet%22%3A%22horizontal%22%2C%22mobile%22%3A%22horizontal%22%7D\" id=\"tsdqcqfGPuEI_{{placeholder content='e3sgcmFuZG9tX2lkIH19' key='menu'}}\" class=\"brz-menu brz-menu__preview brz-menu__mmenu brz-menu--has-dropdown brz-css-n91ed brz-css-1akgap1\"><ul class=\"brz-menu__ul\">{{placeholder content='e3sgbWVudV9sb29wIH19' menuId='dc8d678facf2eea55abec859b6f8c26e3216a0b9a216aa60a6a8b5756e9ea4b0'}}<li data-menu-item-id=\"{{placeholder content='e3sgbWVudV9pdGVtX3VpZCB9fQ=='}}\" class=\"brz-menu__item {{placeholder content='e3sgbWVudV9pdGVtX2NsYXNzbmFtZSB9fQ=='}}\"><a class=\"brz-a\" target=\"{{placeholder content='e3sgbWVudV9pdGVtX3RhcmdldCB9fQ=='}}\" href=\"{{placeholder content='e3sgbWVudV9pdGVtX2hyZWYgfX0='}}\" title=\"{{placeholder content='e3sgbWVudV9pdGVtX2F0dHJfdGl0bGUgfX0='}}\">{{placeholder content='e3sgbWVudV9pdGVtX2ljb24gfX0='}}<span class=\"brz-span\">{{placeholder content='e3sgbWVudV9pdGVtX3RpdGxlIH19'}}</span></a>{{placeholder content='e3sgbWVnYV9tZW51IH19'}}{{placeholder content='e3sgbWVudV9sb29wX3N1Ym1lbnUgfX0='}}<ul class=\"brz-menu__sub-menu\">{{placeholder content='e3sgbWVudV9sb29wIH19' recursive='1'}}{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul>{{placeholder content='e3sgZW5kX21lbnVfbG9vcF9zdWJtZW51IH19'}}</li>{{placeholder content='e3sgZW5kX21lbnVfbG9vcCB9fQ=='}}</ul></nav>{{placeholder content='e3sgZW5kX2dyb3VwIH19'}}</div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-16qa2c6 brz-css-1rcpqcz brz-wrapper\"><div class=\"brz-posts brz-css-cl6qmp brz-css-11ckoy3\" data-brz-custom-id=\"kyxD0eOo0haq\"><div class=\"brz-posts__wrapper\">{{placeholder content='e3ticml6eV9kY19wb3N0X2xvb3B9fQ==' type='posts' collection_type='/collection_types/4959140' count='3' order_by='id' order='DESC' offset='0'}}<div class=\"brz-posts__item\"><div class=\"brz-columns brz-css-9izhcb brz-css-j2nle7 brz-css-tx3g8z\" data-brz-custom-id=\"klpJKINEfaiO\"><div class=\"brz-column__items brz-css-qbnzgb brz-css-rhzuxa\"><div id=\"\" class=\"brz-css-rcv95 brz-css-skwex9 brz-wrapper\"><div class=\"brz-image brz-css-15xpw96 brz-css-17lbhdj\" data-brz-custom-id=\"qmF9ME46aQeI\"><picture class=\"brz-picture brz-d-block brz-p-relative brz-css-gxdnbn brz-css-1ttms7k\"><source srcset=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='263' cW='350'}} 1x, {{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='526' cW='700'}} 2x\" media=\"(min-width: 992px)\"><source srcset=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='266' cW='354'}} 1x, {{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='532' cW='708'}} 2x\" media=\"(min-width: 768px)\"><img class=\"brz-img\" srcset=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='323' cW='430'}} 1x, {{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='646' cW='860'}} 2x\" src=\"{{placeholder content='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ==' cH='263' cW='350'}}\" alt=\"{{placeholder content='e3sgYnJpenlfZGNfaW1hZ2VfYWx0IH19' entityId='' entityType='' imagePlaceholder='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ=='}}\" title=\"{{placeholder content='e3sgYnJpenlfZGNfaW1hZ2VfdGl0bGUgfX0=' entityId='' entityType='' imagePlaceholder='e3ticml6eV9kY19pbWdfZmVhdHVyZWRfaW1hZ2V9fQ=='}}\" draggable=\"false\" loading=\"lazy\"></picture></div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-i7zumc brz-wrapper\"><div class=\"brz-wp-title brz-css-1f4xfm3 brz-css-1slhfhe brz-css-1jcy4c3\" data-brz-custom-id=\"yflMCbIEgZj7\"><span class=\"brz-wp-title-content\" style=\"min-height:20px\">{{placeholder content='e3ticml6eV9kY19wb3N0X3RpdGxlfX0='}}</span></div></div><div class=\"brz-wrapper-clone brz-flex-xs-wrap brz-css-kwv3lv brz-css-wckscm\" data-brz-custom-id=\"kTYuI3BOrgPm\"><a class=\"brz-a brz-btn brz-css-uk96b3 brz-css-z6u47w brz-css-tdg8w brz-css-yrfdqg brz-css-e6b431\" href=\"{{placeholder content='e3ticml6eV9kY191cmxfcG9zdH19'}}\" target=\"_self\" rel=\"noopener\" data-brz-link-type=\"external\" data-brz-custom-id=\"hPQBrYiaVrgJ\"><svg class=\"brz-icon-svg align-[initial] brz-css-w3n9bt brz-css-jxjvb7\"><use href=\"{{brizy_dc_url_site}}/icon/22459456/glyph/tail-right.svg#nc_icon\"></use></svg><span data-brz-translate-text=\"1\" class=\"brz-span brz-text__editor\">READ MORE</span></a></div></div></div></div>{{placeholder content='e3tlbmRfYnJpenlfZGNfcG9zdF9sb29wfX0='}}</div></div></div><div id=\"\" class=\"brz-css-rcv95 brz-css-1oph5n4 brz-wrapper\"><div class=\"brz-post-info brz-css-1k7ppv3 brz-css-1endcla\" data-brz-custom-id=\"l9Pef9gQL8QE\"><div class=\"brz-ui-ed-list brz-ui-ed-list-split brz-ui-ed-post-info\"><div class=\"brz-ui-ed-spin-nested-loading\"><div class=\"brz-ui-ed-spin-container\"><ul class=\"brz-ui-ed-list-items\"><li class=\"brz-ui-ed-list-item brz-ui-ed-list-item-no-flex\"><svg class=\"brz-ui-ed-icon-svg\" width=\"1em\" height=\"1em\" viewBox=\"0 0 16 16\"><path fill=\"currentColor\" d=\"m13.985 11.439-3.063-.77c-.505-.126-.887-.491-.989-.944l-.187-.831c1.493-.6 2.452-1.91 2.454-3.354V3.85C12.235 1.815 10.44.118 8.129.002c-1.136-.03-2.238.345-3.054 1.04C4.26 1.738 3.8 2.695 3.8 3.694V5.54c.001 1.444.96 2.755 2.454 3.355l-.187.83c-.102.452-.484.817-.989.944l-3.063.769C1.414 11.589 1 12.072 1 12.622v2.764c0 .34.313.615.7.615h12.6c.387 0 .7-.275.7-.615v-2.764c0-.55-.414-1.032-1.015-1.183Z\"></path></svg>{{placeholder content='e3tjb2xsZWN0aW9uX2l0ZW1fYXV0aG9yfX0='}}</li><li class=\"brz-ui-ed-list-item brz-ui-ed-list-item-no-flex\"><svg class=\"brz-ui-ed-icon-svg\" width=\"1em\" height=\"1em\" viewBox=\"0 0 16 16\" fill=\"currentColor\"><g stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\"><path d=\"M16,3.5 L16,14.5 C16,15.3284505 15.2325149,16 14.2857143,16 L1.71428571,16 C0.767485119,16 0,15.3284505 0,14.5 L0,3.5 C0,2.67154948 0.767485119,2 1.71428571,2 L3.42857143,2 L3.42857143,0.375 C3.42857143,0.16796875 3.62053571,0 3.85714286,0 L5.28571429,0 C5.52232143,0 5.71428571,0.16796875 5.71428571,0.375 L5.71428571,2 L10.2857143,2 L10.2857143,0.375 C10.2857143,0.16796875 10.4776786,0 10.7142857,0 L12.1428571,0 C12.3794643,0 12.5714286,0.16796875 12.5714286,0.375 L12.5714286,2 L14.2857143,2 C15.2325149,2 16,2.67154948 16,3.5 Z M14,13.8223684 L14,5 L2,5 L2,13.8223684 C2,13.9204359 2.09161932,14 2.20454545,14 L13.7954545,14 C13.9083807,14 14,13.9204359 14,13.8223684 Z\" fill=\"currentColor\" fill-rule=\"nonzero\"></path></g></svg>{{placeholder content='e3tjb2xsZWN0aW9uX2l0ZW1fZGF0ZX19'}}</li><li class=\"brz-ui-ed-list-item brz-ui-ed-list-item-no-flex\"><svg viewBox=\"0 0 16 16\" width=\"1em\" height=\"1em\"><g fill=\"currentColor\"><path fill=\"currentColor\" d=\"M8,0C3.6,0,0,3.6,0,8s3.6,8,8,8s8-3.6,8-8S12.4,0,8,0z M12,9H7V4h2v3h3V9z\"></path></g></svg>{{placeholder content='e3tjb2xsZWN0aW9uX2l0ZW1fdGltZX19'}}</li></ul></div></div></div></div></div></div></div></section>{{end_A}}"

	group := placeholder.New()
	if err := group.AddPlaceholder(&placeholder.Placeholder{
		Name:    "A",
		Label:   "Label A",
		Group:   "Group 1",
		Display: placeholder.DisplayInline,
		Handler: func(node *model.Node) (string, error) {
			time.Sleep(2 * time.Millisecond)
			return fmt.Sprintf("[Value_Block_A_%d]", rand.Intn(1000)), nil
		},
		FallbackHandler: func(node *model.Node) (string, error) {
			return "Fallback_A", nil
		},
	}); err != nil {
		fmt.Errorf("add placeholder failed: %s", err)
	}

	if err := group.AddPlaceholder(&placeholder.Placeholder{
		Name:    "B",
		Label:   "Label B",
		Group:   "Group 2",
		Display: placeholder.DisplayInline,
		Handler: func(node *model.Node) (string, error) {
			time.Sleep(2 * time.Millisecond)
			return fmt.Sprintf("[Value_Block_B_%d]", rand.Intn(1000)), nil
		},
		FallbackHandler: func(node *model.Node) (string, error) {
			return "Fallback_Value_Block_B_", nil
		},
	}); err != nil {
		fmt.Errorf("add placeholder failed: %s", err)
	}
	if err := group.AddPlaceholder(&placeholder.Placeholder{
		Name:    "C",
		Label:   "Label C",
		Group:   "Group 2",
		Display: placeholder.DisplayInline,
		Handler: func(node *model.Node) (string, error) {
			time.Sleep(2 * time.Millisecond)
			return fmt.Sprintf("[Value_C_%d]", rand.Intn(1000)), nil
		},
		FallbackHandler: func(node *model.Node) (string, error) {
			return "Fallback_Value_Block_C_", nil
		},
	}); err != nil {
		fmt.Errorf("add placeholder failed: %s", err)
	}
	if err := group.AddPlaceholder(&placeholder.Placeholder{
		Name:    "placeholder",
		Label:   "Label placeholder",
		Group:   "Group 2",
		Display: placeholder.DisplayInline,
		Handler: func(node *model.Node) (string, error) {
			m := node.Value.(*model.Placeholder)
			decodedBytes, err := base64.StdEncoding.DecodeString(strings.Trim(m.Attrs[0].Value, "'"))
			time.Sleep(50 * time.Millisecond)
			if err != nil {
				return "", err
			}
			return string(decodedBytes), nil
		},
		FallbackHandler: func(node *model.Node) (string, error) {
			return "Fallback_Value_Block_C_", nil
		},
	}); err != nil {
		fmt.Errorf("add placeholder failed: %s", err)
	}

	r := replacer.New(group)
	start := time.Now()
	r.Prepare(input)
	content := r.Replace()
	elapsed := time.Since(start)
	fmt.Printf("Replaced in paralel %s\n", elapsed)
	fmt.Println(content)

	r2 := replacer.New(group)
	start2 := time.Now()
	r2.Prepare(input)
	content = r2.ReplaceSync()
	elapsed2 := time.Since(start2)
	fmt.Printf("Replaced syncron %s\n", elapsed2)
	fmt.Println(content)

	defer fmt.Printf("All Done")
}
