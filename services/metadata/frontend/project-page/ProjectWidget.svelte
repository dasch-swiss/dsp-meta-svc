<script lang='ts'>
  import { projectMetadata } from "../store";
  import type {Text, URL, Person, Organization, Grant, Project} from "../interfaces";

  function findObjectById(id: string): Grant | Person | Organization {
    let grants = $projectMetadata?.grants;
    let g = grants.find(o => o.__id === id);
    if (g) return g;

    let persons = $projectMetadata?.persons
    if (persons && persons.length > 0){
      let p = persons.find(o => o.__id === id);
      if (p) return p;
    }

    let o = $projectMetadata?.organizations.find(o => o.__id === id);
    if (o) return o;
  };

  const truncateString = (s: string) => {
    if (s.length > 35) {
      return `${s.slice(0, 35)}...`;
    } else return s;
  };

  function getText(text: Text, lang?:string) {
    if (!text){
      return ""
    }

    let langs = Object.keys(text);
    
    if (langs.length === 0) {
      return ""
    } else if (lang && langs.includes(lang)) {
      return text[lang]
    } else if (langs.includes('en')) {
      return text['en']
    } else {
      return text[langs[0]]
    }
  }

</script>

<div class=label>DSP Internal Shortcode</div>
<div class=data>{$projectMetadata?.project.shortcode}</div>

<div class=label>Data Management Plan</div>
<div class=data>{$projectMetadata?.project.dataManagementPlan ? "available" : "unavailable"}</div>

<div class=label>Discipline</div>
<!-- TODO: remove array check. should be done by interface? -->
{#if Array.isArray($projectMetadata?.project.disciplines)}
  {#each $projectMetadata?.project.disciplines as d}
    {#if d.__type === "URL"}
      <a class="data external-link" href={d.url} target=_>{truncateString(d.text)}</a>
    {:else}
      {#if getText(d).match(/^\d+ /)}
        <a class="data external-link" href=http://www.snf.ch/SiteCollectionDocuments/allg_disziplinenliste.pdf target=_>{truncateString(getText(d))}</a>
      {:else}
        <div class="data">{getText(d)}</div>
      {/if}
    {/if}
  {/each}
{/if}

<div class=label>Temporal Coverage</div>
{#if Array.isArray($projectMetadata?.project.temporalCoverage)}
  {#each $projectMetadata?.project.temporalCoverage as t}
    {#if t.__type === "URL"}
      <a class="data external-link" href={t.url} target=_>{truncateString(t.text)}</a>
    {:else}
      <div class="data">{getText(t)}</div>
      <!-- <div class="data">{getText(asText(t))}</div> -->
    {/if}
  {/each}
{/if}

<div class=label>Spatial Coverage</div>
{#if Array.isArray($projectMetadata?.project.spatialCoverage)}
  {#each $projectMetadata?.project.spatialCoverage as s}
    <a class="data external-link" style="text-transform: capitalize" href={s.url} target=_>{truncateString(s.text)}</a>
  {/each}
{/if}

<div class=label>Start date</div>
<div class=data>{$projectMetadata?.project.startDate}</div>

{#if $projectMetadata?.project.endDate}
<div class=label>End date</div>
<div class=data>{$projectMetadata?.project.endDate}</div>
{/if}


<div class=label>Funder</div>
{#if Array.isArray($projectMetadata?.project.funders)}
  {#each $projectMetadata?.project.funders.map((o) => {return findObjectById(o)}) as f}
    {#if f.__type === "Person"}
      {console.log('person',f)}
      <!-- TODO: handle funding person - need to find example -->
      <!-- <div class=data>{findObjectById(f)?.givenName.split(";").join(" ")} {findObjectById(f)?.familyName}</div> -->
      {:else if f.__type === "Organization"}
      <div class=data>{f.name}</div>
    {/if}
  {/each}
{/if}
  
{#if $projectMetadata?.project.grants && Array.isArray($projectMetadata?.project.grants)}
  <div class=label>Grant</div>
  {#each $projectMetadata?.project.grants.map(id => {return findObjectById(id)}) as g}
    {#if g.__type === "Grant"}
      {#if g?.number && g?.url && g?.name}
        <a class="data external-link" href={g?.url.url} target=_>{truncateString(`${g?.number}: ${g?.name}`)}</a>
        <!-- TODO: roll back if people don't like it -->
        <!-- <a class="data external-link" href={g?.url.url} target=_>{g?.number}</a> -->
      {:else if g?.number && g?.url}
        <a class="data external-link" href={g?.url.url} target=_>{g?.number}</a>
      {:else if g?.number}
        <span class="data">{g?.number}</span>
      {:else}
        {#each [g?.funders[0]].map(o => {return findObjectById(o)}) as f}
          {#if f.__type === "Organization"}
            <span class="data">{f.name}</span>
          {/if}
        {/each}
      {/if}
    {/if}
  {/each}
{/if}

<!-- TODO -->
{#if $projectMetadata?.project.contactPoint}
  <div class=label>Contact</div>
  {#each [findObjectById($projectMetadata?.project.contactPoint)] as c}
    {#if c.__type === 'Organization'}
      <div id=contact class=data>{c.name}</div>
      {#if c.email}
        <a class="data email" href="mailto:{c?.email}">{c?.email}</a>
      {/if}
    {:else if c.__type === 'Person'}
      {#if c?.givenNames && c?.familyNames}
        <div id=contact class=data>{c?.givenNames?.join(" ")} {c?.familyNames.join(" ")}</div>
      {/if}
      {#if Array.isArray(c?.affiliation)}
        {#each c?.affiliation as o}
          {#each [findObjectById(o)] as org}
            {#if org.__type === 'Organization'}
              <span class="data">{org.name}</span>
            {/if}
          {/each}
        {/each}
      {/if}
      {#if c.emails}
        <a class="data email" href="mailto:{c?.emails[0]}">{c?.emails[0]}</a>
      {/if}
    {/if}
  {/each}
{/if}

<div class=label>Project Website</div>
{#if Array.isArray($projectMetadata?.project.urls)}
  {#each $projectMetadata?.project.urls as url}
    <a class="data external-link" href={url.url} target=_>{truncateString(url.text)}</a>
  {/each}
{/if}

{#if $projectMetadata?.project}
  <div class=label>Keywords</div>
  <span class="keyword">{$projectMetadata?.project.keywords.map(t => {return getText(t)}).join(", ")}</span>
{/if}

<style>
  a {
    display: block;
    color: var(--lead-colour);
  }
  .keyword {
    padding: 0;
    /* display: inline;
    border: 1px solid #cdcdcd;
    border-radius: 0.25rem;
    color: #fff;
    background-color: var(--third);
		box-shadow: var(--shadow-1);
    white-space: pre;
    line-height: 2em;
    padding: 4px; */
  }
  /* .keyword:hover {
    color: var(--third);
    background-color: #fff;
    border-color: var(--third);
  } */
  .label, .data {
    margin: 5px 0;
  }
  .label {
    padding: 10px 0 0;
  }
</style>
