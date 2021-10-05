<script lang="ts">
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import { handleSnackbar } from "../store";
  import { getText, findOrganizationByID, findObjectByID } from "../functions";
  import type { Dataset } from "../interfaces";

  export let dataset: Dataset;

  let isAbstractExpanded: boolean;
  let abstractLinesNumber: number;

  const toggleExpand = () => {
    isAbstractExpanded = !isAbstractExpanded;
  };

  onMount(() => {
    const el = document.getElementById('abstract');
    const lineHeight = parseInt(window.getComputedStyle(el).getPropertyValue('line-height'));
    const divHeight = el.scrollHeight;
    abstractLinesNumber = divHeight / lineHeight;
    isAbstractExpanded = abstractLinesNumber > 6 ? false : true;
  });

  const copyToClipboard = () => {
    let text = document.createRange();
    text.selectNode(document.getElementById('how-to-cite'));
    window.getSelection().removeAllRanges();
    window.getSelection().addRange(text);
    document.execCommand('copy');
    window.getSelection().removeAllRanges();
    handleSnackbar.set({isSnackbar: true, message: 'Citation copied succesfully!'});
  };

  const truncateString = (s: string) => {
    const browserWidth = window.innerWidth;
    if (browserWidth < 992 && s.length > ((browserWidth - 100) / 8)) {
      return `${s.substring(0, (browserWidth - 100) / 8)}...`;
    } else if (browserWidth >= 992 && s.length > (browserWidth / 17)) {
      return `${s.substring(0, (browserWidth / 17))}...`;
    } else return s;
  };
</script>

<div id=dataset in:fade={{duration: 200}}>
  {#if dataset}
  <!-- Alternative titles -->
    {#if dataset?.alternativeTitles}
      <div>
        <span class=label>Alternative Title</span>
        <span class=data>{dataset?.alternativeTitles.map((t => {return getText(t)})).join(', ')}</span>
      </div>
    {/if}

    <div class=grid-wrapper>
      <!-- Access conditions -->
      <div>
        <span class=label>Access</span>
          {#if dataset?.accessConditions}
            <span class=data>{dataset?.accessConditions}</span>
          {:else}
            <span class="warning data">access conditions missing</span>
          {/if}
      </div>

      <!-- Status -->
      <div>
        <span class=label>Status</span>
          {#if dataset?.status}
            <span class=data>{dataset?.status}</span>
          {:else}
            <span class="warning data">status missing</span>
          {/if}
      </div>

      <!-- Dates -->
      {#if dataset.dateCreated}
        <div>
          <span class=label>Date Created</span>
          <span class=data>{dataset?.dateCreated}</span>
        </div>
      {/if}
      {#if dataset.datePublished}
        <div>
          <span class=label>Date Published</span>
          <span class=data>{dataset?.datePublished}</span>
        </div>
      {/if}
      {#if dataset.dateModified}
        <div>
          <span class=label>Date Modified</span>
          <span class=data>{dataset?.dateModified}</span>
        </div>
      {/if}

      <!-- Type of Data -->
      <div>
        <span class=label>Type of Data</span>
          {#if dataset?.typeOfData}
            <span class=data>{dataset?.typeOfData.join(', ')}</span>
          {:else}
            <span class="warning data">type of data missing</span>
          {/if}
      </div>

      <!-- Additional -->
      {#if dataset?.additional}
        <div style="grid-column-start: 1;grid-column-end: 3;">
          <span class=label>Additional documentation</span>
          {#each dataset?.additional as d}
            {#if d.__type === "URL"}
              <a class="data external-link" href={d.url} target=_>{truncateString(d.text)}</a>
            {:else}
              <span class=data>{getText(d)}</span>
            {/if}
          {/each}
        </div>
      {/if}
      
    </div>

    <!-- License -->
    <!-- TODO: check how this looks with multiple licenses -->
    <div>
      <span class=label>License</span>
        {#if dataset?.licenses}
          {#each dataset?.licenses as l}
            <div class=data>
              <a href={l.license.url} class=external-link target=_>{l.license.text}</a>
              {#if l.details}
                <div>{l.details}</div>
              {/if}
              <div>({l.date})</div>
            </div>
          {/each}
        {:else}
          <span class="warning data">licenses missing</span>
        {/if}
    </div>

    <!-- Languages -->
    <div class=grid-wrapper style="grid-template-columns: repeat(1, 1fr)">
      <div>
        <span class=label>Languages</span>
        {#if dataset?.languages}
          <span class=data>{dataset?.languages.map(l => {return getText(l)}).join(', ')}</span>
        {:else}
          <span class="warning data">languages missing</span>
        {/if}
      </div>
    </div>

    <!-- URLs -->
    {#if dataset?.urls}
      <div class=grid-wrapper style="grid-template-columns: repeat(1, 1fr)">
        <div>
          <span class=label>Dataset Website</span>
          {#each dataset?.urls as u}
            {#if u.__type === 'URL'}
              <div><a class="data external-link" href={u.url} target=_>{truncateString(u.text)}</a></div>
            {/if}
          {/each}
        </div>
      </div>
    {/if}

    <!-- How To Cite -->
    <div class="property-row">
      <span class=label style="display:inline">
        How To Cite
        {#if dataset?.howToCite}
          <button on:click={copyToClipboard} title="copy citation to the clipboard">
            <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>
          </button>
        {/if}
      </span>
      {#if dataset?.howToCite}
        <span id=how-to-cite class=data>{dataset?.howToCite}</span>
      {:else}
        <span class="warning data">how to cite missing</span>
      {/if}
    </div>

    <!-- Abstract -->
    <div>
      <span class=label>Abstract</span>
      {#if dataset?.abstracts}
        <div id=abstract class="data {isAbstractExpanded ? '' : 'abstract-short'}">
          {#each dataset?.abstracts as a}
            {#if a.__type === "URL"}
              <div><a class="data external-link" href={a.url} target=_>{truncateString(a.text)}</a></div>
            {:else}
              <div>{getText(a)}</div>
            {/if}
          {/each}
        </div>
      {:else}
        <span class="warning data" id="abstract">abstract missing</span>
      {/if}
    </div>
    {#if abstractLinesNumber > 6}
      <div on:click={toggleExpand} class=expand-button>show {isAbstractExpanded ? "less" : "more"}</div>
    {/if}

    <!-- Attribution -->
    <span class=label>Attributions</span>
    <div class=grid-wrapper>
      {#if dataset?.attributions}
        {#each dataset?.attributions as a}
          <div class="attributions data">
            <div class=role>{a.roles.join(", ")}</div>
            {#each [findObjectByID(a.agent)] as p}
              {#if p.__type === 'Person'}
                {#if p.authorityRefs}
                  <a href={p.authorityRefs[0].url} target=_ class="external-link">{p.givenNames.join(" ")} {p.familyNames.join(" ")}</a>
                {:else}
                  <div>{p.givenNames.join(" ")} {p.familyNames.join(" ")}</div>
                {/if}
                {#if p.affiliation}
                  {#each p.affiliation.map(o => {return findOrganizationByID(o)}) as org}
                    <div>{org.name}</div>
                  {/each}
                {/if}
                <div>{p.jobTitles[0]}</div>
                {#if p.email}
                  <a class=email href="mailto:{p.email}">{p.email}</a>
                {/if}
              {:else if p.__type === 'Organization'}
                {#if p.url}
                  <a href={p.url.url} target=_ class="external-link">{p.name}</a>
                {/if}
                {#if p.email}
                  <a class=email href="mailto:{p.email}">{p.email}</a>
                {/if}
              {/if}
            {/each}
          </div>
        {/each}
      {:else}
        <span class="warning data">attributions missing</span>
      {/if}
    </div>

  {/if}
</div>

<style>
  a {color: var(--lead-colour);}
  button {
    border: none;
    background-color: inherit;
    padding: 0;
    position: relative;
    top: 10px;
    color: var(--lead-colour);
    z-index: 0;
  }
  .icon {
    margin: -1rem 0 0.25rem;
  }
  .icon:hover {
    color: var(--dasch-light-violet);
  }
  .role {
    color: var(--secondary-colour);
  }
  .abstract-short {
    display: -webkit-box;
    -webkit-line-clamp: 6;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
  .attributions {
    padding: 10px 10px 0 0;
    line-height: 1.5;
  }
  .grid-wrapper {
    display: grid;
    grid-template-columns: repeat(1, 1fr);
  }
  @media screen and (min-width: 576px) {
    .grid-wrapper {
      grid-template-columns: repeat(2, 1fr);
    }
  }
  @media screen and (min-width: 1200px) {
    .grid-wrapper {
      grid-template-columns: repeat(3, 1fr);
    }
  }
</style>
