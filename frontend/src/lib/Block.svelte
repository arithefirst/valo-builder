<script lang="ts">
    import Skin from './Skin.svelte'
    export let weapon: string;
    export let visible: boolean;

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

    let src: string
    let uuid: string
    weapons[weapon].subscribe((val) => {
        src = val.src
        uuid = val.uuid
    })

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/skin/${weapon}`);
        return await response.json();
    }

    async function fetchChroma() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/chromas?uuid=${uuid}`);
        return await response.json();
    }

    let chromaResponse = fetchChroma();
    function refetchChroma() {
        chromaResponse = fetchChroma();
    }

    function toggleVis() {
        visible = false;
    }

</script>

<div class="blur"/>
<div class="skins-selector">
    <svg xmlns="http://www.w3.org/2000/svg" class="close" viewBox="0 0 512 512"
         on:click={toggleVis}
         on:keypress={toggleVis}
         role="button"
         tabindex=0>
        <!--!Font Awesome Free 6.6.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.-->
        <path d="M256 48a208 208 0 1 1 0 416 208 208 0 1 1 0-416zm0 464A256 256 0 1 0 256 0a256 256 0 1 0 0 512z
        M175 175c-9.4 9.4-9.4 24.6 0 33.9l47 47-47 47c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l47-47 47 47
        c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9l-47-47 47-47c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0l-47
        47-47-47c-9.4-9.4-24.6-9.4-33.9 0z" fill="#6c7086"/>
    </svg>
    <h2>{weapon.charAt(0).toUpperCase() + weapon.slice(1)}</h2>
    <div class="levels-img-chroma">
        {#await chromaResponse then chromaData}
            <img src={chromaData.chromas[0].fullRender} alt={chromaData.chromas[0].displayName}/>
        {/await}
    </div>
    <hr>
    <div class="scrollable">
        <div class="grid">
            {#await fetchData() then data}
                {#each data as skin, i}
                    <!-- Skin component requires i var for tabindex -->
                    <Skin {i} {weapon} src={skin.fullRender} title={skin.displayName} uuid={skin.uuid} refreshFunc={refetchChroma}></Skin>
                {/each}
            {/await}
        </div>
    </div>
</div>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        padding-left: 5px;
    }

    .blur {
        position: fixed;
        content: '';
        backdrop-filter: blur(10px);
        width: 100vw;
        height: 100vh;
        top: 0;
        left: 0;
    }

    .skins-selector {
        border: 3px solid #6c7086;
        background-color: #1e1e2e;
        border-radius: 15px;
        width: 818px;
        height: calc(100vh - 26px);
        text-align: center;
        z-index: 1000;
        position: absolute;
        margin: 10px 50% 10px;
        translate: -50%;
        overflow: hidden;
    }

    .scrollable {
        height: calc(100% - 250px);
        position: absolute;
        overflow-y: scroll;
    }

    .close {
        position: fixed;
        right: 10px;
        top: 10px;
        width: 4%;
        cursor: pointer;
        overflow: hidden;
    }

    .levels-img-chroma {
        height: 150px;
    }

    img {
        height: 80%;
    }

    hr {
        width: 90%;
        height: 2px;
        border-color: #6c7086;
        background-color: #6c7086;
    }

</style>
