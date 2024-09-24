<script>
    export let src;
    export let alt
    export let i;

    import {card} from './stores.js'

    function load() {
        // Lookup the currently set weapon in the weapons table;
        // Then set that weapon's store's value to that of src
        card.set( {src: src, name: alt} )
    }
</script>

<div
    on:click={load}
    on:keyup={load}
    role="button"
    tabindex={i}
>
    <p><strong>{alt}</strong></p>
    <img {src} {alt}>
</div>

<style>

    p {
        color: #cdd6f4;
    }

    div {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        border: 3px solid #6c7086;
        width: 180px;
        background-color: #1e1e2e;
        text-align: center;
        font-size: 0.8rem;
        border-radius: 15px;
        margin: 10px 5px;
        cursor: pointer;
    }

    img {
        width: 90%;
        padding: 10px;
    }

</style>