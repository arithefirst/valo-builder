<script lang="ts">
    import Skin from './Skin.svelte'
    import Buddy from './Buddy.svelte'
    export let weapon: string;
    export let visible: boolean;

    let buddyMode: boolean = false;
    function modeSwitcher(mode: boolean) {
        buddyMode = mode;
        const buddy = document.getElementById("buddies-button");
        const skin = document.getElementById("skins-button");
        if (buddyMode) {
            buddy.classList.add("active")
            skin.classList.remove("active");
        } else {
            skin.classList.add("active")
            buddy.classList.remove("active");
        }
    }

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
    let buddy: string
    weapons[weapon].subscribe((val) => {
        href = val.src
        uuid = val.uuid
        buddy = val.buddy
    })

    let buddies = fetchBuddies();
    let resp = fetchData();
    let buddyQuery: string;
    let query: string;

    // Function for handling searches for buddies
    async function buddySearch() {
        if (buddyQuery === "") {
            // If the query is empty, give the default data instead
            console.log("Search empty. Re-setting to default.")
            buddies = fetchBuddies();
        } else {
            const response = await fetch(`http://127.0.0.1:8080/api/v1/search?q=${buddyQuery}&weapon=20`)
            buddies = await response.json()
            console.log(buddies)
        }
    }

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

    async function fetchBuddies() {
        const response = await fetch("http://127.0.0.1:8080/api/v1/buddy");
        return await response.json();
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

<div id="active-skin-cover"></div>
<div id="active-buddy-cover"></div>
<div class="mode-selector">
    <div class="active" id="skins-button" on:click={() => {modeSwitcher(false)}} on:keypress={() => {modeSwitcher(false)}} role="button" tabindex=0>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path fill="#e2e2e4" d="M7 5h16v4h-1v1h-6a1 1 0 0 0-1 1v1
            a2 2 0 0 1-2 2H9.62c-.38 0-.73.22-.9.56l-2.45 4.89c-.17.34-.51.55-.89.55
            H2s-3 0 1-6c0 0 3-4-1-4V5h1l.5-1h3zm7 7v-1a1 1 0 0 0-1-1h-1s-1 1 0 2
            a2 2 0 0 1-2-2a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1"/>
        </svg>
    </div>
    <div id="buddies-button" on:click={() => {modeSwitcher(true)}} on:keypress={() => {modeSwitcher(true)}} role="button" tabindex=0>
        <svg id="Layer_1" data-name="Layer 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 122.5 122.88"><defs>
            <style>.cls-1{fill-rule:evenodd;}</style></defs><title>keychain</title><path fill="#e2e2e4" class="cls-1" d="M9,9A30.63,30.63,0,0,1,30.7,0h
            0A30.73,30.73,0,0,1,59.05,42.52l4.37,4.37,1-1a10.06,10.06,0,0,1,14.21,0l41,41a10.09,10.09,0,0,1,0,14.22l-18.84,18.84a10,10,0,0,1-7.11,2.93h
            0a10,10,0,0,1-7.11-2.94l-41-41a10.09,10.09,0,0,1,0-14.22l1-1L42,59.26A30.71,30.71,0,0,1,9,9ZM67.1,50.58,72,55.47a1.71,1.71,0,0,1,0,2.39L57.48,72.38
            a1.7,1.7,0,0,1-2.38,0l-4.9-4.9-1,1a4.88,4.88,0,0,0,0,6.84l41,41a4.82,4.82,0,0,0,3.42,1.4h0A4.86,4.86,0,0,0,97,116.26l18.84-18.85
            a4.85,4.85,0,0,0,0-6.84l-41-41a4.85,4.85,0,0,0-6.84,0l-1,1ZM37.88,55.16l-3.35-3.35a1.68,1.68,0,0,1,0-2.38L49.05,34.91a1.68,1.68,0,0,1,2.38,0
            L55,38.46a25.66,25.66,0,0,0,1.21-7.76A25.5,25.5,0,0,0,30.7,5.21h0a25.49,25.49,0,0,0,0,51,25.75,25.75,0,0,0,7.18-1Zm16.63-.27
            a4.34,4.34,0,1,1,0,6.14,4.33,4.33,0,0,1,0-6.14Z"/></svg>
    </div>
</div>
<div class="blur" on:click={toggleVis} on:keypress={toggleVis} role="button" tabindex=0 />
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
    {#if !buddyMode}
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
                        {#each data as skin}
                            <!-- Skin component requires i var for tabindex -->
                            <Skin {weapon} src={skin.fullRender} title={skin.displayName} uuid={skin.uuid} bind:chromaIndex={chromaIndex} refreshFunc={refetchChroma}></Skin>
                        {/each}
                    {/if}
                {/await}
            </div>
        </div>
        <!-- Skin mode only styles -->
        <style>
            #active-buddy-cover {
                visibility: hidden;
            }
        </style>
    {/if}
    {#if buddyMode}
        <h2>Buddies</h2>
        {#if (buddy !== "")}
            <img src={buddy} alt="Currently equipped buddy" style="height: 20%; margin-bottom: 30px;"/>
        {/if}
        <div class="search">
            <input
                placeholder="Search..."
                bind:value={buddyQuery}
                on:input={buddySearch}
            />
        </div>
        <hr>
        <div class="buddy-scrollable">
            <div class="buddy-grid">
                {#await buddies then data}
                    {#if data !== null}
                        {#each data as buddy}
                            <Buddy src={buddy.displayIcon} title={buddy.displayName} {weapon} />
                        {/each}
                    {/if}
                {/await}
            </div>
        </div>
        <!-- Buddy mode only styles -->
        <style>
            #active-skin-cover {
                visibility: hidden;
            }
        </style>
    {/if}
</div>

{#if weapon === "melee"}
    <style>
        #active-skin-cover,
        #active-buddy-cover,
        .mode-selector {
            visibility: hidden;
        }
    </style>
{/if}

<style>
    input {
        transform: translate(-50%, 0);
        border-radius: 7px;
        position: absolute;
        bottom: 0;
        left: 50%;
        /* Colors */
        background-color: #1e1e2e;
        color: #cdd6f4;
        border-color: #6c7086;
    }

    .grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        padding-left: 5px;
    }

    .buddy-grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
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
        z-index: 1;
        position: fixed;
        content: '';
        backdrop-filter: blur(15px);
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
        z-index: 3;
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

    .buddy-scrollable {
        height: calc(100% - 140px);
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

    .search {
        position: relative;
        height: 10px;
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

    .mode-selector {
        height: 100%;
        width: 100%;
        position: absolute;
        top: 30px;
        left: calc(50% + 399px);

    }

    #skins-button,
    #buddies-button {
        cursor: pointer;
        border-radius: 5px;
        padding-left: 10px;
        margin-bottom: 7px;
        line-height: 45px;
        text-align: center;
        position: relative;
        z-index: 2;
        background-color: #181825;
        border: 3px solid #6c7086;
        height: 45px;
        width: 30px;
    }

    svg {
        display: flex;
        transform: translate(2px, 10px);
        justify-self: center;
        align-self: center;
        width: 80%;
    }

    #active-skin-cover {
        width: 3px;
        position: absolute;
        height: 44px;
        background-color: #1e1e2e;
        left: calc(50% + 409px);
        top: 33px;
        z-index: 4;
    }

    #active-buddy-cover {
        width: 3px;
        position: absolute;
        height: 44px;
        background-color: #1e1e2e;
        left: calc(50% + 409px);
        top: 91px;
        z-index: 4;
    }

    .active {
        background-color: #1e1e2e !important;
    }
</style>