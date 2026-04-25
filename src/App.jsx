import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import heroImg from './assets/hero.png'
import './App.css'
import Footer from './components/Footer/Footer'
import Header from './components/Header/Header'
import Homepage from './components/Homepage/Homepage'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Header />

      <Homepage />
      
      <Footer />
    </>

  )
}

export default App
