/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      maxWidth: {
        "container": "1200px"
      },
      flex: {
        "3/1": "3",
        "1/3": "calc(33.333% - 0.5rem) 0 0",
        "logo": "12rem 0 0"
      },
    },
  },
  plugins: [],
};
