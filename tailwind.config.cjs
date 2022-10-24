const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}"],
  darkMode: "class", // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        primary: colors.blue,
        secondary: colors.purple,
        gray: colors.stone,
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  // daisyUI config
  daisyui: {
    styled: true,
    themes: true,
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    prefix: "",
    darkTheme: "dark",
  },
};
