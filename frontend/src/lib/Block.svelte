<script lang="ts">
    import Skin from './Skin.svelte'
    export let weapon: string;
    export let visible: boolean;


    // Function to hide the swatches if there is only 1
    function hideSwatch(json: object) {
        if (Object.keys(json).length === 1) {
            return "display: none;"
        } else return ""
    }

    // Import stores from stores.js
    import { classic, shorty, frenzy, ghost,
        sheriff, stinger, spectre, bucky,
        judge, bulldog, guardian, phantom,
        vandal, marshal, outlaw, operator,
        ares, odin, melee} from './stores.js'

    // Mapping of each weapon for each store
    const weapons = {
        "classic": classic, "shorty": shorty, "frenzy": frenzy,
        "ghost": ghost, "sheriff": sheriff, "stinger": stinger,
        "spectre": spectre, "bucky": bucky, "judge": judge,
        "bulldog": bulldog, "guardian": guardian, "phantom": phantom,
        "vandal": vandal, "marshal": marshal, "operator": operator,
        "ares": ares, "odin": odin, "melee": melee, "outlaw": outlaw
    }

    // Mapping of each weapon to it's index
    const indices = {
        "odin":0, "ares":1, "vandal":2, "bulldog":3,
        "phantom":4, "judge":5, "bucky":6, "frenzy":7,
        "classic":8, "ghost":9, "sheriff":10, "shorty":11,
        "operator":12, "guardian":13, "outlaw":14, "marshal":15,
        "spectre":16, "stinger":17, "melee":18
    }

    // Which chroma to use
    let chromaIndex = 0

    // Subscribe to the store for the current weapon
    let href: string
    let uuid: string
    weapons[weapon].subscribe((val) => {
        href = val.src
        uuid = val.uuid
    })


    let resp = fetchData();
    let query: string;

    // Function for handling searches
    async function search() {
        if (query === "") {
            // If the query is empty, give the default data instead
            console.log("Search empty. Re-setting to default.")
            resp = fetchData();
        } else {
            const response = await fetch(`http://127.0.0.1:8080/api/v1/search?q=${query}&weapon=${indices[weapon]}`)
            resp = await response.json()
        }
    }

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/skin/${weapon}`);
        return await response.json();
    }

    async function fetchChroma() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/chromas?uuid=${uuid}`);
        return await response.json();
    }

    // Function to re-fetch data when skin changes
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
    <svg xmlns="http://www.w3.org/2000/svg" class="close" viewBox="0 0 24 24"
         on:click={toggleVis}
         on:keypress={toggleVis}
         role="button"
         tabindex=0>
         <!-- Close Icon -->
         <path fill="#6c7086" d="
         M18.3 5.71a.996.996 0 0 0-1.41 0L12 10.59
         L7.11 5.7A.996.996 0 1 0 5.7 7.11L10.59 12
         L5.7 16.89a.996.996 0 1 0 1.41 1.41
         L12 13.41l4.89 4.89a.996.996 0 1 0 1.41-1.41
         L13.41 12l4.89-4.89c.38-.38.38-1.02 0-1.4"/>
    </svg>
    <h2>{weapon.charAt(0).toUpperCase() + weapon.slice(1)}</h2>
    <div class="levels-img-chroma">
        {#await chromaResponse then chromaData}
            <div class="chromas">
                {#each chromaData.chromas as chroma, i}
                    <!-- svelte-ignore a11y-invalid-attribute permalink -->
                    <a
                            href="#"
                            role="button"
                            tabindex={i}
                            on:click={() => {
                            chromaIndex = i
                            weapons[weapon].set( {src: chromaData.chromas[chromaIndex].fullRender, uuid: uuid} )
                        }}
                            on:keypress={() => {
                            chromaIndex = i
                            weapons[weapon].set( {src: chromaData.chromas[chromaIndex].fullRender, uuid: uuid} )
                        }}
                    >
                        <img
                                class="swatch"
                                id="swatch-{i}"
                                src={chroma.swatch}
                                alt={chroma.displayName}
                                style="{hideSwatch(chromaData.chromas)}"
                        />
                    </a>
                {/each}
            </div>
            <div>
                <img id="weapon" src={href} alt={chromaData.chromas[chromaIndex].displayName}/>
                <input
                    id="search"
                    placeholder="Search..."
                    bind:value={query}
                    on:input={search}
                />
            </div>
        {/await}
    </div>
    <hr>
    <div class="scrollable">
        <div class="grid">
            {#await resp then data}
                {#if data !== null}
                    {#each data as skin, i}
                        <!-- Skin component requires i var for tabindex -->
                        <Skin {i} {weapon} src={skin.fullRender} title={skin.displayName} uuid={skin.uuid} bind:chromaIndex={chromaIndex} refreshFunc={refetchChroma}></Skin>
                    {/each}
                {/if}
            {/await}
        </div>
    </div>
</div>

<style>
    input {
        transform: translate(-50%, 0);
        border-radius: 7px;
        position: absolute;
        bottom: 0;
        left: 50%;
    }

    .grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        padding-left: 5px;
    }

    .chromas {
        position: absolute;
        left: 0;
        bottom: 0;
        margin-top: 10px;
        margin-left: 40px;
        display: inline-block;
        height: 40px;
        width: 180px
    }

    .swatch {
        cursor: pointer;
        border-radius: 3px;
        height: 40px;
        width: 40px;
    }

    a {
        height: 40px;
        width: 40px;
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
        height: calc(100% - 300px);
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
        position: relative;
        height: 200px;
    }

    #weapon {
        transform: translate(-50%, 0);
        position: absolute;
        top: 20px;
        left: 50%;
        height: 65%;
    }

    hr {
        width: 90%;
        height: 2px;
        border-color: #6c7086;
        background-color: #6c7086;
    }

</style>
