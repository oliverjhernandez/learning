/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {},
    fontFamily: {
      sans: ['ui-sans-serif', 'system-ui'],
      serif: ['ui-serif', 'Georgia'],
      mono: ['ui-monospace', 'SFMono-Regular'],
      lvl1: ['Pacifico'],
      lvl2: ['Plus Jakarta Sans'],
      body: ['"Open Sans"'],
    },
    backgroundColor: {
      primary: '#ededed',
      seconday: '#FFFFFF',
    },
  },
  plugins: [],
}
