/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  safelist: [
      'rotate-1','rotate-2', 'rotate-3', 'rotate-4', 'rotate-5', 'rotate-6',
      '-rotate-1', '-rotate-2', '-rotate-3', '-rotate-4', '-rotate-5', '-rotate-6',
      'ml-1', 'ml-2', 'ml-3', 'ml-4', 'ml-5', 'ml-6', 'ml-7', 'ml-8', 'ml-9', 'ml-10',
      '-ml-1', '-ml-2', '-ml-3', '-ml-4', '-ml-5', '-ml-6', '-ml-7', '-ml-8', '-ml-9', '-ml-10'
  ],
  theme: {
    fontFamily: {
      'sans': ['circe', 'ui-sans-serif', 'helvetica']
    },
    extend: {
      rotate: {
        '1': '2deg',
        '2': '4deg',
        '3': '7deg',
        '4': '11deg',
        '5': '14deg',
        '7': '18deg',
        '8': '22deg',
        '9': '24deg',
      }
    }
  },
  plugins: []
};