import './app.css';
import logo from "./assets/images/frückteulen.papa.transparent.png";
import { Greet } from "../wailsjs/go/main/App";
import { useState } from "preact/hooks";

export function App(props) {
    const [resultText, setResultText] = useState("Bitte gib deinen Namen ein 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <>
            <div id="App">
                <img src={logo} id="logo" alt="logo" />
                <div id="result">{resultText}</div>
                <br />
                <input id="name" onChange={updateName} autoComplete="off" name="input" type="text" />
                <br />
                <br />
                <button class="green small" onClick={greet}>Hey 👋</button>&nbsp;
                <button class="yellow small" onClick={greet}>Hey 👋</button>&nbsp;
                <button class="red small" onClick={greet}>Hey 👋</button>&nbsp;
                <button class="blue small" onClick={greet}>Hey 👋</button>&nbsp;
                <button class="orange small" onClick={greet}>Hey 👋</button>&nbsp;
                <button class="lila small" onClick={greet}>Hey 👋</button>&nbsp;
                <button class="small" onClick={greet}>Hey 👋</button>
            </div >
        </>
    )
}
