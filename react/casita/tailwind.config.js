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
      primary: '#FCFBFC',
      secondary: '#FFFFFF',
      blue: '#4169E1',
      labelgreen: '#E9FFE5',
    },
    borderColor: {
      labelgreen: '#92C585',
      card: '#D1D5DB',
    },
    textColor: {
      logo: '#F87171',
    },
  },
  plugins: [],
}
