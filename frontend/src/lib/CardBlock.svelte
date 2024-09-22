<script lang="ts">
    import Card from './Card.svelte'
    export let visible: boolean;

    let resp = fetchData();
    let query: string;

    // Function for handling searches
    async function search() {
        if (query === "") {
            // If the query is empty, give the default data instead
            console.log("Search empty. Re-setting to default.")
            resp = fetchData();
        } else {
            const response = await fetch(`http://127.0.0.1:8080/api/v1/search?q=${query}&weapon=19`)
            resp = await response.json()
        }
    }

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/cards`);
        return await response.json();
    }

    function toggleVis() {
        visible = false;
    }

</script>

<div class="blur"/>
<div class="cards-selector">
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
    <h2>Cards</h2>
    <hr>
    <div class="scrollable">
        <div class="grid">
            {#await resp then data}
                {#if data !== null}
                    {#each data as card, i}
                        <!-- Card component requires i var for tabindex -->
                        <Card {i} alt={card.displayName} src={card.largeArt}></Card>
                    {/each}
                {/if}
            {/await}
        </div>
    </div>
</div>

<style>

    .grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        padding-left: 10px;
    }

    .blur {
        z-index: 20;
        position: fixed;
        content: '';
        backdrop-filter: blur(15px);
        width: 100vw;
        height: 100vh;
        top: 0;
        left: 0;
    }

    .cards-selector {
        border: 3px solid #6c7086;
        background-color: #1e1e2e;
        border-radius: 15px;
        width: 800px;
        height: calc(100vh - 26px);
        text-align: center;
        z-index: 1000;
        position: absolute;
        margin: 10px 50% 10px;
        translate: -50%;
        overflow: hidden;
    }

    .scrollable {
        height: 100%;
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

    hr {
        width: 90%;
        height: 2px;
        border-color: #6c7086;
        background-color: #6c7086;
    }

</style>