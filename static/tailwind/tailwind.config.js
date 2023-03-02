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
  daisyui: {
    themes: ["light", "dark", "cupcake", "bumblebee", "emerald", "corporate", "synthwave", "retro", "cyberpunk", "valentine", "halloween", "garden", "forest", "aqua", "lofi", "pastel", "fantasy", "wireframe", "black", "luxury", "dracula", "cmyk", "autumn", "business", "acid", "lemonade", "night", "coffee", "winter"],
  },
  plugins: [require("daisyui")],
}
