import './app.css';
import logo from "./assets/images/frückteulen.papa.transparent.png";
import { Color } from "../wailsjs/go/main/App";
import { Greet } from "../wailsjs/go/main/App";
import { ChoosColor } from "../wailsjs/go/main/App";
import { useState } from "preact/hooks";

export function First(props) {
    // useState('') gibt ein Array zurück: [aktuellerWert, setFunktion].
    // const [<VARIABLE>, set<VARIABLE-Functionsaufruf>] = useState('<STARTWERT>');
    const [resultColor, setResultColor] = useState('');
    const [resultText, setResultText] = useState('Bitte gib deinen Namen ein 👇');
    const [name, setName] = useState('');

    // function updateName(e) {
    //     setName(e.target.value);
    // }
    // Arrow-Function:
    const updateName = (e) => setName(e.target.value);
    // function updateResultColor(result) {
    //     setResultColor(result);
    // }
    // Arrow-Function:
    const updateResultColor = (result) => setResultColor(result);
    // function updateResultText(result) {
    //     setResultText(result);
    // }
    // Arrow-Function:
    const updateResultText = (result) => setResultText(result);

    function greet(color) {
        Color(color).then(updateResultColor);
        Greet(name).then(updateResultText);
    }

    return (
        <>
            <div id="First">
                <img src={logo} id="logo" alt="logo" />
                <font class={resultColor}>{resultText}</font>
                <br />
                <br />
                <input onChange={updateName} autoComplete="off" name="input" type="text" />
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

export function Second(props) {
    const [resultColor, setResultColor] = useState('');
    const [resultText, setResultText] = useState('Bitte wähle deine Farbe 👇');

    const updateResultColor = (result) => setResultColor(result);
    const updateResultText = (result) => setResultText(result);

    function color(color) {
        Color(color).then(updateResultColor);
        ChoosColor(color).then(updateResultText);
    }

    return (
        <>
            <div id="Second">
                <font class={resultColor}>{resultText}</font>
                <br />
                <br />
                <input class="smallest" onChange={() => color('green')} id="green" name="color" type="radio" /><label class="green" for="green">green</label>
                <input class="smallest" onChange={() => color('yellow')} id="yellow" name="color" type="radio" /><label class="yellow" for="yellow">yellow</label>
                <input class="smallest" onChange={() => color('red')} id="red" name="color" type="radio" /><label class="red" for="red">red</label>
                <input class="smallest" onChange={() => color('blue')} id="blue" name="color" type="radio" /><label class="blue" for="blue">blue</label>
                <input class="smallest" onChange={() => color('orange')} id="orange" name="color" type="radio" /><label class="orange" for="orange">orange</label>
                <input class="smallest" onChange={() => color('lila')} id="lila" name="color" type="radio" /><label class="lila" for="lila">lila</label>
                <input class="smallest" onChange={() => color('"no"')} id="no" name="color" type="radio" /><label class="" for="no">"no"</label>
            </div >
        </>
    );
}