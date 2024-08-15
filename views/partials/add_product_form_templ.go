// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AddProductForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"add-product-form\" class=\"flex flex-col\" hx-post=\"/admin/products/add\" action=\"/admin/products/add\" hx-encoding=\"multipart/form-data\" x-data=\"{\n      productName: undefined,\n      productPrice: undefined,\n      productStockQuantity: 1,\n    }\"><div class=\"overflow-y-auto flex flex-col md:flex-row w-full mt-4\"><div class=\"flex-col md:mr-8 flex flex-1\"><div class=\"mb-5\"><label for=\"ProductName\" class=\"block text-xs font-medium text-gray-700\">Product Name </label> <input x-model.lazy=\"productName\" name=\"productName\" type=\"text\" id=\"ProductName\" placeholder=\"Product Name\" class=\"mt-1 w-full rounded-md border-gray-200 shadow-sm sm:text-sm\"> <span x-show=\"productName &amp;&amp; productName.length &lt; 3\" class=\"text-red-500\">Enter a valid product name</span></div><div class=\"flex flex-col sm:flex-row sm:space-x-2 w-full\"><div class=\"mb-5 sm:w-1/2\"><label for=\"ProductPrice\" class=\"block text-xs font-medium text-gray-700\">Price </label> <input x-model.number=\"productPrice\" name=\"procutPrice\" type=\"number\" id=\"ProductPrice\" placeholder=\"Price\" class=\"h-10 w-full rounded-md border-gray-200 shadow-sm sm:text-sm\"></div><div class=\"m-0 mb-5 sm:w-1/2\"><label for=\"StockQuantity\" class=\"block text-xs font-medium text-gray-700\">Stock Quantity </label><div class=\"flex items-center rounded-md shadow-sm border border-gray-200\"><button x-on:click=\"\n              if (productStockQuantity == 0) {\n                return;\n              }\n              productStockQuantity--\n              \" type=\"button\" class=\"px-3 lg:px-2 size-10 leading-10 text-gray-600 transition hover:opacity-75\">&minus;</button> <input x-model.number=\"productStockQuantity\" name=\"productQuantity\" type=\"number\" id=\"StockQuantity\" value=\"0\" class=\"h-10 w-full border-transparent text-center [-moz-appearance:_textfield] sm:text-sm [&amp;::-webkit-inner-spin-button]:m-0 [&amp;::-webkit-inner-spin-button]:appearance-none [&amp;::-webkit-outer-spin-button]:m-0 [&amp;::-webkit-outer-spin-button]:appearance-none\"> <button x-on:click=\"\n              productStockQuantity++\n              \" type=\"button\" class=\"px-3 lg:px-2 size-10 leading-10 text-gray-600 transition hover:opacity-75\">&plus;</button></div></div></div><div class=\"mb-5 flex flex-col\" x-data=\"{ imageUrls: [], url: &#39;&#39; }\"><label for=\"ImageUrl\" class=\"block text-xs font-medium text-gray-700\">Image URL</label><div class=\"flex\"><input x-model=\"url\" type=\"text\" id=\"ImageUrl\" placeholder=\"URL\" class=\"mt-1 w-full rounded-md border-gray-200 shadow-sm sm:text-sm\"><div x-on:click=\"\n            if (!url || url.length == 0) {\n              return;\n            }\n            if (imageUrls.includes(url)) {\n              url = &#39;&#39;;\n              return;\n            }\n            imageUrls.push(url);\n            url = &#39;&#39;;\n            \" class=\"cursor-pointer ml-1 flex items-center justify-center border p-1 rounded-md px-5 text-white bg-blue-500\"><span class=\"select-none\">Add</span></div><div x-show:=\"imageUrls.length &gt; 0\" x-on:click=\"imageUrls = []\" class=\"cursor-pointer ml-1 flex items-center justify-center border p-1 rounded-md px-5 text-white bg-gray-500\"><span>Clear</span></div></div><ul class=\"grid grid-cols-3 gap-1 mt-2\"><template x-for=\"imageUrl in imageUrls\"><li class=\"py-1 px-1 border rounded-md bg-gray-300 border-gray-400\"><img :src=\"imageUrl\" class=\"rounded-md object-cover\" :alt=\"imageUrl.split(&#39;/&#39;).last\"><div class=\"flex flex-row justify-end items-end\"><span x-on:click=\"imageUrls = imageUrls.filter(item =&gt; item != imageUrl)\" class=\"mt-1 cursor-pointer text-red-500 text-xs font-bold\">DELETE</span></div></li></template></ul></div></div><div class=\"flex flex-col flex-1\"><div class=\"flex flex-col\" x-data=\"{ productImages: [] }\"><p class=\"text-gray-700 text-xs\">Upload product images</p><label for=\"file-input\"><div class=\"cursor-pointer hover:shadow-sm flex-col mt-1 p-10 border-gray-200 rounded-md border bg-gray-100 flex items-center justify-center\"><div class=\"w-20 h-30\"><svg viewBox=\"0 0 21 21\" xmlns=\"http://www.w3.org/2000/svg\"><g fill=\"none\" fill-rule=\"evenodd\" stroke=\"#000000\" stroke-linecap=\"round\" stroke-linejoin=\"round\" transform=\"translate(4 3)\"><path d=\"m8.5 14.5h2c1.1045695 0 2-.8954305 2-2v-8l-4-4h-6c-1.1045695 0-2 .8954305-2 2v10c0 1.1045695.8954305 2 2 2h2\"></path> <path d=\"m3.5 7.5 3-3 3 3\"></path> <path d=\"m6.5 4.5v11\"></path></g></svg></div><span class=\"text-gray-700\">Select product image</span> <input accept=\"image/*\" multiple=\"true\" id=\"file-input\" name=\"productImageFiles\" type=\"file\" class=\"hidden text-center w-full text-sm text-slate-500\n      file:py-2 file:px-4\n      file:rounded-full file:border-0\n      file:text-sm file:font-semibold\n      file:bg-violet-50 file:text-violet-700\n      hover:file:bg-violet-100\" x-on:change=\"\n                  const files = $event.target.files;\n                  for (let i = 0; i &lt; files.length; i++) {\n                    productImages.push(files[i]);\n                  }\n                \"></div></label><div x-show=\"productImages.length != 0\" class=\"grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-2 mt-2\"><template x-for=\"image in productImages\"><div class=\"border-blue-300 border rounded\"><img class=\"object-cover w-full h-20 \" :src=\"URL.createObjectURL(image)\" alt=\"\"></div></template></div><div x-on:click=\"productImages = []\" x-show=\"productImages.length != 0\" class=\"cursor-pointer bg-gray-400 rounded-md p-1 mt-2 flex justify-center items-center\"><p class=\"text-white font-bold\">CLEAR</p></div></div></div></div><div class=\"mt-4 self-end\"><button class=\"inline-block w-full rounded-lg bg-green-500 px-5 py-3 font-medium text-white sm:w-auto\" form=\"add-product-form\">Submit</button></div></form><span hx-swap-oob=\"true\" id=\"path-text\" class=\"py-4 text-sm text-gray-500\"><a href=\"/admin/products\" hx-target=\"#main-section\" hx-get=\"/admin/products\">Product Section</a> > Add Product</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
