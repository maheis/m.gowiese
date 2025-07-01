import './app.css';
import logo from "./assets/images/frückteulen.papa.transparent.png";
import { Color } from "../wailsjs/go/main/App";
import { Greet } from "../wailsjs/go/main/App";
import { useState } from "preact/hooks";

export function App(props) {
    const [resultColor, setResultColor] = useState('asdf');
    const [resultText, setResultText] = useState('Bitte gib deinen Namen ein 👇');
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultColor = (result) => setResultColor(result);
    const updateResultText = (result) => setResultText(result);

    function greet(color) {
        Greet(name).then(updateResultText);
        Color(color).then(updateResultColor);
    }

    return (
        <>
            <div id="App">
                <img src={logo} id="logo" alt="logo" />
                <font class={resultColor}>{resultText}</font>
                <br />
                <br />
                <input id="name" onChange={updateName} autoComplete="off" name="input" type="text" />
                <br />
                <br />
                <button class="green small" onClick={() => greet('green')}>Hey 👋</button>&nbsp;
                <button class="yellow small" onClick={() => greet('yellow')}>Hey 👋</button>&nbsp;
                <button class="red small" onClick={() => greet('red')}>Hey 👋</button>&nbsp;
                <button class="blue small" onClick={() => greet('blue')}>Hey 👋</button>&nbsp;
                <button class="orange small" onClick={() => greet('orange')}>Hey 👋</button>&nbsp;
                <button class="lila small" onClick={() => greet('lila')}>Hey 👋</button>&nbsp;
                <button class="small" onClick={() => greet('')}>Hey 👋</button>
            </div >
        </>
    )
}
