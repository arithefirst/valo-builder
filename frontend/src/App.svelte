<script lang="ts">
  import Item from './lib/Item.svelte';
  import Block from "./lib/Block.svelte";

  // Load all the stores from stores.js
  import { classic } from './lib/stores.js'
  import { shorty } from './lib/stores.js'
  import { frenzy } from './lib/stores.js'
  import { ghost } from './lib/stores.js'
  import { sheriff } from './lib/stores.js'
  import { stinger } from './lib/stores.js'
  import { spectre } from './lib/stores.js'
  import { bucky } from './lib/stores.js'
  import { judge } from './lib/stores.js'
  import { bulldog } from './lib/stores.js'
  import { guardian } from './lib/stores.js'
  import { phantom } from './lib/stores.js'
  import { vandal } from './lib/stores.js'
  import { marshal } from './lib/stores.js'
  import { operator } from './lib/stores.js'
  import { ares } from './lib/stores.js'
  import { odin } from './lib/stores.js'
  import { melee } from './lib/stores.js'
  import { outlaw } from './lib/stores.js'

  // Func and vars for creating the <Block />
  let visible = false;
  let weapon: string;
  function toggleVis(t: string) {
      visible = !visible;
      weapon = t
  }
</script>

<div class="weaponsGrid">
    <div class="sidearms">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("classic")}} on:click={() => {toggleVis("classic")}}><Item weapon="classic" href="{$classic}" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("shorty")}} on:click={() => {toggleVis("shorty")}}><Item weapon="shorty" href="{$shorty}" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("frenzy")}} on:click={() => {toggleVis("frenzy")}}><Item weapon="frenzy" href="{$frenzy}" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("ghost")}} on:click={() => {toggleVis("ghost")}}><Item weapon="ghost" href="{$ghost}" type="pistol"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("sheriff")}} on:click={() => {toggleVis("sheriff")}}><Item weapon="sheriff" href="{$shorty}" type="pistol"/></div>
    </div>

    <div class="smgs-shotguns">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("stinger")}} on:click={() => {toggleVis("stinger")}}><Item weapon="stinger" href="{$shorty}" type="smg"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("spectre")}} on:click={() => {toggleVis("spectre")}}><Item weapon="spectre" href="{$shorty}" type="smg"/></div>
        <p/>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("bucky")}} on:click={() => {toggleVis("bucky")}}><Item weapon="bucky" href="{$shorty}" type="shotgun"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("judge")}} on:click={() => {toggleVis("judge")}}><Item weapon="judge" href="{$shorty}" type="shotgun"/></div>
    </div>

    <div class="rifles-melee">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("bulldog")}} on:click={() => {toggleVis("bulldog")}}><Item weapon="bulldog" href={$bulldog} type="rifle"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("guardian")}} on:click={() => {toggleVis("guardian")}}><Item weapon="guardian" href={$guardian} type="rifle"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("phantom")}} on:click={() => {toggleVis("phantom")}}><Item weapon="phantom" href={$phantom} type="rifle"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("vandal")}} on:click={() => {toggleVis("vandal")}}><Item weapon="vandal" href={$vandal} type="rifle"/></div>
        <p/>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("melee")}} on:click={() => {toggleVis("melee")}}><Item weapon="melee" href={$melee} type="knife"/></div>
    </div>

    <div class="snipers-lmgs">
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("marshal")}} on:click={() => {toggleVis("marshal")}}><Item weapon="marshal" href={$marshal} type="sniper"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("outlaw")}} on:click={() => {toggleVis("outlaw")}}><Item weapon="outlaw" href={$outlaw} type="sniper"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("operator")}} on:click={() => {toggleVis("operator")}}><Item weapon="operator" href={$operator} type="sniper"/></div>
        <p/>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("ares")}} on:click={() => {toggleVis("ares")}}><Item weapon="ares" href={$ares} type="lmg"/></div>
        <div role="button" tabindex=0 on:keypress={() => {toggleVis("odin")}} on:click={() => {toggleVis("odin")}}><Item weapon="odin" href={$odin} type="lmg"/></div>
    </div>
</div>

<footer>RIOT GAMES, VALORANT and any associated logos are trademarks, service marks, and/or registered trademarks of <a href="https://riotgames.com">Riot Games, Inc.</a></footer>

<style>
    footer {
        font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
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