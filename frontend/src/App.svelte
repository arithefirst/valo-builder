<script lang="ts">

  import Item from './lib/Item.svelte';
  import Block from "./lib/Block.svelte";
  import {card} from "./lib/stores.js"
  import CardBlock from "./lib/CardBlock.svelte"

  let cardSrc: string;
  let cardAlt: string;
  card.subscribe((val) => {
      cardSrc = val.src
      cardAlt = val.name
  })

 // Func and vars for creating the <Block /> and <CardBlock />
  let visible = false;
  let cardVisible = false;
  let weapon: string;
  function toggleVis(t: string) {
          visible = !visible
          weapon = t
  }
  function toggleCardVis() {
      cardVisible = !cardVisible
  }
</script>

{#if visible}
    <Block {weapon} bind:visible={visible}/>
{/if}
{#if cardVisible}
    <CardBlock bind:visible={cardVisible}/>
{/if}

<div class="weaponsGrid">
    <div class="sidearms">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("classic")}} on:click={() => {toggleVis("classic")}}><Item weapon="classic" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("shorty")}} on:click={() => {toggleVis("shorty")}}><Item weapon="shorty" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("frenzy")}} on:click={() => {toggleVis("frenzy")}}><Item weapon="frenzy" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("ghost")}} on:click={() => {toggleVis("ghost")}}><Item weapon="ghost" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("sheriff")}} on:click={() => {toggleVis("sheriff")}}><Item weapon="sheriff" type="pistol"/></div>
    </div>

    <div class="smgs-shotguns">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("stinger")}} on:click={() => {toggleVis("stinger")}}><Item weapon="stinger" type="smg"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("spectre")}} on:click={() => {toggleVis("spectre")}}><Item weapon="spectre" type="smg"/></div>
        <p/>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("bucky")}} on:click={() => {toggleVis("bucky")}}><Item weapon="bucky" type="shotgun"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("judge")}} on:click={() => {toggleVis("judge")}}><Item weapon="judge" type="shotgun"/></div>
    </div>

    <div class="rifles-melee">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("bulldog")}} on:click={() => {toggleVis("bulldog")}}><Item weapon="bulldog" type="rifle"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("guardian")}} on:click={() => {toggleVis("guardian")}}><Item weapon="guardian" type="rifle"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("phantom")}} on:click={() => {toggleVis("phantom")}}><Item weapon="phantom" type="rifle"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("vandal")}} on:click={() => {toggleVis("vandal")}}><Item weapon="vandal" type="rifle"/></div>
        <p/>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("melee")}} on:click={() => {toggleVis("melee")}}><Item weapon="melee" type="knife"/></div>
    </div>

    <div class="snipers-lmgs">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("marshal")}} on:click={() => {toggleVis("marshal")}}><Item weapon="marshal" type="sniper"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("outlaw")}} on:click={() => {toggleVis("outlaw")}}><Item weapon="outlaw" type="sniper"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("operator")}} on:click={() => {toggleVis("operator")}}><Item weapon="operator" type="sniper"/></div>
        <p/>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("ares")}} on:click={() => {toggleVis("ares")}}><Item weapon="ares" type="lmg"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("odin")}} on:click={() => {toggleVis("odin")}}><Item weapon="odin" type="lmg"/></div>
    </div>
    <div class="card">
        <!-- svelte-ignore a11y-no-noninteractive-element-to-interactive-role -->
        <img
            src={cardSrc}
            alt={cardAlt}
            role="button"
            tabindex=0
            on:keypress={toggleCardVis}
            on:click={toggleCardVis}
        >
    </div>
</div>

<footer>RIOT GAMES, VALORANT and any associated logos are trademarks, service marks, and/or registered trademarks of <a href="https://riotgames.com/">Riot Games, Inc.</a></footer>

<style>
    footer {
        font-family: sans-serif;
        position: absolute;
        text-align: center;
        padding-top: 10px;
        color: white;
        width: 100%;
        bottom: 0;
    }

    a {
        color: inherit;
        text-decoration: underline;
    }

    .card {
        position: relative;
        margin: 2.75330396476vw auto 0;
        height: 30.160857908vw;
        cursor: pointer;
        mask-image: linear-gradient(0deg, rgba(255,255,255,0) 0%, rgba(255,255,255,1) 40%);
    }

    .card img {
        height: 100%;
    }

    .weaponsGrid{display: flex;}

    .sidearms {
        margin-top: 2.75330396476vw;
        margin-left: 4.40528634361vw;
        width: 16.519823788vw;
    }

    .smgs-shotguns, .rifles-melee, .snipers-lmgs{
        margin-top: 2.75330396476vw;
        margin-left: 2.20264317181vw;
        width: 16.519823788vw;
    }

    p {
        margin-bottom: 3.74449339207vw
    }

    :root {
        overflow: hidden;
        background-attachment: fixed;
        background-image: url("./assets/bg_main.png");
    }
</style>