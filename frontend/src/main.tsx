import { render } from 'preact'
import { App } from './App'
import 'picocss/css/pico.min.css'

render(<App />, document.getElementById('app-root') as HTMLElement)
