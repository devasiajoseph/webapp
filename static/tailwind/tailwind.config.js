/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../../html/**/*.{html,js}","../../cljs/src/**/*.cljs"],
  darkMode: 'class',
  theme: {
    container: {
      center: true,
    },
    extend: {},
  },
  plugins: [require("daisyui")],
}
