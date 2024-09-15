<script>
    import Skin from './Skin.svelte'
    export let weapon;

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/skin/${weapon}`);
        return await response.json();
    }
</script>

<div class="skins-selector">
    <h2>{weapon.charAt(0).toUpperCase() + weapon.slice(1)}</h2>
    <hr>
    <div class="grid">
        {#await fetchData() then data}
            {#each data as skin, i}
                <!-- Skin component requires i var for tabindex -->
                <Skin {i} {weapon} src={skin.fullRender} title={skin.displayName}></Skin>
            {/each}
        {/await}
    </div>
</div>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        padding-left: 5px;
    }

    .skins-selector {
        border-radius: 15px;
        border: 3px solid #6c7086;
        width: 818px;
        height: calc(100vh - 20px);
        text-align: center;
        z-index: 1000;
        position: absolute;
        margin: 10px 50%;
        translate: -50%;
        overflow: scroll;

    }

    hr {
        width: 90%;
        height: 2px;
        border-color: #6c7086;
        background-color: #6c7086;
    }

</style>
