import { render } from 'preact';
import { First } from './app';
import { Second } from './app';
import './style.css';

render(<First />, document.getElementById('first'));
render(<Second />, document.getElementById('second'));