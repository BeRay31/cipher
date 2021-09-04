import { defineConfig } from 'windicss/helpers'

export default defineConfig({
  darkMode: 'class',
  // https://windicss.org/posts/v30.html#attributify-mode
  attributify: true,
  theme: {
    extend: {
      fontFamily: {
        poppins: ['Poppins', 'sans-serif'],
        spartan: ['Spartan', 'sans-serif']
      },
      textColor: {
        primary: '#E8E8E8',
        secondary: '#66739D',
        alternate: '#A8AAB0'
      },
      backgroundColor: {
        primary: '#2D313D',
        secondary: '#353744',
        highlight: '#66739D'
      },
      borderColor: theme => ({
        ...theme('colors'),
        primary: '#66739D'
      })
    }
  }
})
