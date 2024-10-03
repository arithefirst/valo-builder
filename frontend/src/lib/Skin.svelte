<script lang="ts">
    export let src;
    export let title;
    export let chromaIndex
    export let refreshFunc;
    export let weapon;
    export let uuid;

    // Import stores from stores.js
    import { classic, shorty, frenzy, ghost,
        sheriff, stinger, spectre, bucky,
        judge, bulldog, guardian, phantom,
        vandal, marshal, outlaw, operator,
        ares, odin, melee} from './stores.js'

    const weapons = {
        "classic": classic, "shorty": shorty, "frenzy": frenzy,
        "ghost": ghost, "sheriff": sheriff, "stinger": stinger,
        "spectre": spectre, "bucky": bucky, "judge": judge,
        "bulldog": bulldog, "guardian": guardian, "phantom": phantom,
        "vandal": vandal, "marshal": marshal, "operator": operator,
        "ares": ares, "odin": odin, "melee": melee, "outlaw": outlaw
    }

    function load() {
        // Lookup the currently set weapon in the weapons table;
        // Then set that weapon's store's value to that of src
        let buddy: string;
        weapons[weapon].subscribe((val) => {buddy = val.buddy})
        weapons[weapon].set( {src: src, uuid: uuid, buddy: buddy} )
        chromaIndex = 0
        refreshFunc()
    }
</script>

<div
    class="skinCon"
    on:click={load}
    on:keyup={load}
    role="button"
    tabindex={0}
>
    <p><strong>{title}</strong></p>
    <img {src} alt="{title}">
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
        width: 250px;
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