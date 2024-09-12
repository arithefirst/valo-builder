<script>
    import Skin from './Skin.svelte'
    export let weapon;
    let posts;

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/skin/${weapon}`);
        const data = await response.json();
        console.log(data);
        return data;
    }

    console.log(posts);

</script>


{#await fetchData() then data}
    <div id="outer">
        {#each data as skin}
                <Skin src={skin.fullRender} title={skin.displayName}></Skin>
        {/each}
    </div>
{/await}

<style>
    #outer {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 10px;
        width: 768px;
    }
</style>
