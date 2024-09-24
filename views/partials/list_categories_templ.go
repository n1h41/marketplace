// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "n1h41/marketplace/internal/domain/productdmn"

func ListCategories(categories []productdmn.Category) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"category-list\" class=\"flex flex-col sample-transition md:px-20 xl:px-72 2xl:px-96\"><div class=\"flex mt-5\"><a class=\"inline-block rounded border border-indigo-600 bg-indigo-600 px-12 py-3 text-sm font-medium text-white hover:bg-transparent hover:text-indigo-600 focus:outline-none focus:ring active:text-indigo-500\" hx-get=\"/admin/categories/add\" hx-target=\"#category-list\" hx-swap=\"outerHTML transition:true\">Create</a></div><div></div><div><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range categories {
			if !item.IsSubCategory {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"p-2 border-gray-300 border mt-2 rounded\"><span class=\"text-blue-800 font-semibold\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(item.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/partials/list_categories.templ`, Line: 29, Col: 19}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><div class=\"pl-5\"><ul>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, subItem := range categories {
					if item.Id == subItem.Parent {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><span class=\"text-blue-700 underline\">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var3 string
						templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(subItem.Name)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/partials/list_categories.templ`, Line: 41, Col: 27}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div><span hx-swap-oob=\"true\" id=\"path-text\" class=\"py-4 text-sm text-gray-500\"><a href=\"/admin/products\" hx-swap=\"innerHTML transition:true\" hx-target=\"#main-section\" hx-get=\"/admin/products\">Product</a> > Category</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
