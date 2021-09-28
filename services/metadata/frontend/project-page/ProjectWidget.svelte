<script lang='ts'>
  import { projectMetadata } from "../store";
  import { getText, findObjectByID, findGrantByID, findOrganizationByID } from "../functions";

  const truncateString = (s: string) => {
    // TODO: can this be improved? 1. dynamic langth depending on space; 2. show full text on hover
    if (s.length > 35) {
      return `${s.slice(0, 35)}...`;
    } else return s;
  };

</script>


{#if $projectMetadata}

<!-- Shortcode -->
  <div class=label>DSP Internal Shortcode</div>
  <div class=data>{$projectMetadata?.project.shortcode}</div>

  <!-- DMP -->
  {#if $projectMetadata?.project.dataManagementPlan}
    <div class=label>Data Management Plan</div>
    {#if $projectMetadata?.project.dataManagementPlan.url}
      <a class="data external-link" href="{$projectMetadata?.project.dataManagementPlan.url.url}" target=_>
        {$projectMetadata?.project.dataManagementPlan ? "available" : "unavailable"}
      </a>
    {:else}
      <div class=data>{$projectMetadata?.project.dataManagementPlan ? "available" : "unavailable"}</div>
    {/if}
  {/if}

  <!-- How to Cite -->
  <div class=label>How to Cite this Project</div>
  <div class=data>{$projectMetadata?.project.howToCite}</div>

  <!-- Disciplines -->
  <div class=label>Disciplines</div>
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

  <!-- Temporal Coverage -->
  <div class=label>Temporal Coverage</div>
  {#each $projectMetadata?.project.temporalCoverage as t}
    {#if t.__type === "URL"}
      <a class="data external-link" href={t.url} target=_>{t.text ? truncateString(t.text) : truncateString(t.url)}</a>
    {:else}
      <div class="data">{getText(t)}</div>
    {/if}
  {/each}

  <!-- Spatial Coverage -->
  <div class=label>Spatial Coverage</div>
  {#each $projectMetadata?.project.spatialCoverage as s}
    <a class="data external-link" style="text-transform: capitalize" href={s.url} target=_>{truncateString(s.text)}</a>
  {/each}

  <!-- Start Date -->
  <div class=label>Start date</div>
  <div class=data>{$projectMetadata?.project.startDate}</div>
  
  <!-- End Date -->
  {#if $projectMetadata?.project.endDate}
    <div class=label>End date</div>
    <div class=data>{$projectMetadata?.project.endDate}</div>
  {/if}

  <!-- Funders -->
  <div class=label>Funder</div>
  {#each $projectMetadata?.project.funders.map((o) => {return findObjectByID(o)}) as f}
    {#if f.__type === "Person"}
      {console.log('person',f)}
      <!-- TODO: handle funding person - need to find example -->
      <!-- <div class=data>{findObjectById(f)?.givenName.split(";").join(" ")} {findObjectById(f)?.familyName}</div> -->
    {:else if f.__type === "Organization"}
      <div class=data>{f.name}</div>
    {/if}
  {/each}
  
  <!-- Grants -->
  {#if $projectMetadata?.project.grants}
    <div class=label>Grant</div>
    {#each $projectMetadata?.project.grants.map(id => {return findGrantByID(id)}) as g}
      {#if g?.number && g?.url && g?.name}
        <a class="data external-link" href={g?.url.url} target=_>{truncateString(`${g?.number}: ${g?.name}`)}</a>
        <!-- TODO: roll back if people don't like it -->
        <!-- <a class="data external-link" href={g?.url.url} target=_>{g?.number}</a> -->
      {:else if g?.number && g?.url}
        <a class="data external-link" href={g?.url.url} target=_>{g?.number}</a>
      {:else if g?.number}
        <span class="data">{g?.number}</span>
      {:else}
        {#each [g?.funders[0]].map(o => {return findOrganizationByID(o)}) as f}
          <span class="data">{f.name}</span>
        {/each}
      {/if}
    {/each}
  {/if}

  <!-- Contact Point -->
  {#if $projectMetadata?.project.contactPoint}
    <div class=label>Contact</div>
    {#each [findObjectByID($projectMetadata?.project.contactPoint)] as c}
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
            {#each [findOrganizationByID(o)] as org}
              <span class="data">{org.name}</span>
            {/each}
          {/each}
        {/if}
        {#if c.emails}
          <a class="data email" href="mailto:{c?.emails[0]}">{c?.emails[0]}</a>
        {/if}
      {/if}
    {/each}
  {/if}

  <!-- URLs -->
  <div class=label>Project Website</div>
  <a class="data external-link" href={$projectMetadata?.project.url.url} target=_>{truncateString($projectMetadata?.project.url.text)}</a>
  <!-- Secondary URL -->
  {#if $projectMetadata?.project.secondaryURL}
    <a class="data external-link" href={$projectMetadata?.project.secondaryURL.url} target=_>{truncateString($projectMetadata?.project.secondaryURL.text)}</a>
  {/if}

  <!-- Keywords -->
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
