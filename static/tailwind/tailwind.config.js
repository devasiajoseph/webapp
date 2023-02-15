/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../../html/**/*.{html,js}","../../cljs/src/**/*.{cljs,js}"],
  darkMode: 'class',
  theme: {
    container: {
      center: true,
    },
    extend: {},
  },
  plugins: [],
}
