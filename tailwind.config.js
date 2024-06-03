/** @type {import('tailwindcss').Config} */
module.exports = {
content: [
"./views/**/*.{html,js,templ}",
"./components/**/*.{html,js,templ}",
],
theme: {
extend: {},
},
plugins: [require("@tailwindcss/forms")],
};
