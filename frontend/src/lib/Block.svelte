<script>
    import Skin from './Skin.svelte'
    export let weapon;

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/skin/${weapon}`);
        const data = await response.json();
        console.log(data);
        return data;
    }
</script>

<div class="outer">
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

    .outer {
        border-radius: 15px;
        border: 3px solid #6c7086;
        width: 818px;
        text-align: center;
        margin: 10px auto;

    }

    hr {
        width: 90%;
        height: 2px;
        border-color: #6c7086;
        background-color: #6c7086;
    }

</style>
