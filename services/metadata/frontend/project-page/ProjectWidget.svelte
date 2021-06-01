<script>
  import { currentProject, currentProjectMetadata } from "../store";

  let grant;

  const findObjectById = (id) => {
    
    // console.log('searching', $currentProjectMetadata?.grants);
    let res = $currentProjectMetadata?.grants.find(o => o['@id'] === 'id');
    console.log(res);
    if (res) return res;
    res = $currentProjectMetadata?.persons.find(o => o['@id'] === 'id');
    console.log(res);
    if (res) return res;
    res = $currentProjectMetadata?.organizations.find(o => o['@id'] === 'id');
    console.log(res);
    if (res) return res;
    // console.log('searching', $currentProjectMetadata?.persons);
    // console.log('searching', $currentProjectMetadata?.grants);
    // console.log('searching', $currentProjectMetadata?.find(o => o['@type'] === 'Dataset'));
    // console.log($currentProjectMetadata?.find(obj => obj['@id'] === id));
    return "blah"
    // grant = $currentProjectMetadata?.metadata.find(obj => obj.id === id);
    // return $currentProjectMetadata?.find(obj => obj['@id'] === id);
  };

  // const handleSpatialCoverageName = (s) => {
  //   const regex = /[^/]+\.html/i;
  //   // return s.split("/")[4].split('.')[0].split("-").join(' ');
  //   return s.substr(s.lastIndexOf('/') + 1).split('.')[0].split("-").join(' ');
  //   // return s.match(regex)[0].split('.')[0].split("-").join(' ');
  // }

  const truncateString = (s) => {
    if (s.length > 35) {
      return `${s.slice(0, 35)}...`;
    } else return s;
  };
</script>

<div class=label>DSP Internal Shortcode</div>
<div class=data>{$currentProject?.shortcode}</div>

<div class=label>Data Management Plan</div>
<div class=data>{$currentProject?.dataManagementPlan ? "available" : "unavailable"}</div>

<div class=label>Discipline</div>
{#if Array.isArray($currentProject?.disciplines)}
  {#each $currentProject?.disciplines as d}
    {#if d.url}
      <a class="data external-link" href={d.url} target=_>{truncateString(d.text)}</a>
    {:else}
      {#if d[Object.keys(d)[0]].match(/^\d+ /)}
        <a class="data external-link" href=http://www.snf.ch/SiteCollectionDocuments/allg_disziplinenliste.pdf target=_>{truncateString(d[Object.keys(d)[0]])}</a>
      {:else}
        <div class="data">{d[Object.keys(d)[0]]}</div>
      {/if}
    {/if}
  {/each}
{/if}

<div class=label>Temporal Coverage</div>
{#if Array.isArray($currentProject?.temporalCoverage)}
  {#each $currentProject?.temporalCoverage as t}
    {#if t.url}
      <a class="data external-link" href={t.url} target=_>{truncateString(t.text)}</a>
    {:else}
      <div class="data">{t[Object.keys(t)[0]]}</div>
    {/if}
  {/each}
{/if}

<div class=label>Spatial Coverage</div>
{#if Array.isArray($currentProject?.spatialCoverage)}
  {#each $currentProject?.spatialCoverage as s}
    <a class="data external-link" style="text-transform: capitalize" href={s.url} target=_>{truncateString(s.text)}</a>
  {/each}
{/if}

<div class=label>Start date</div>
<div class=data>{$currentProject?.startDate}</div>

{#if $currentProject?.endDate}
<div class=label>End date</div>
<div class=data>{$currentProject?.endDate}</div>
{/if}

<div class=label>Funder</div>
{#if Array.isArray($currentProject?.funders)}
  {#each $currentProject?.funders as f}
  {console.log(f)}
  {console.log('searching for: ', f, 'found: ',findObjectById(f))}
    <!-- {#if findObjectById(f.id).type === "http://ns.dasch.swiss/repository#Person"}
      <div class=data>{findObjectById(f.id)?.givenName.split(";").join(" ")} {findObjectById(f.id)?.familyName}</div>
    {:else if findObjectById(f.id).type === "http://ns.dasch.swiss/repository#Organization"}
      <div class=data>{findObjectById(f.id)?.name.join(", ")}</div>
    {/if} -->
  {/each}
{/if}

{#if $currentProject?.grant && Array.isArray($currentProject?.grant)}
<div class=label>Grant</div>
  {#each $currentProject?.grant as g}
    {#if findObjectById(g.id)?.number && findObjectById(g.id)?.url}
      <a class="data external-link" href={findObjectById(g.id)?.url[0].url} target=_>{findObjectById(g.id)?.number}</a>
    {:else if findObjectById(g.id)?.number}
      <span class="data">{findObjectById(g.id)?.number}</span>
    {:else}
      <span class="data">{findObjectById(findObjectById(g.id)?.funder[0].id)?.name.join(', ')}</span>
    {/if}
  {/each}
{/if}

{#if $currentProject?.contactPoint}
  <div class=label>Contact</div>
  {#if findObjectById($currentProject?.contactPoint[0].id)?.givenName && findObjectById($currentProject?.contactPoint[0].id)?.familyName}
    <div id=contact class=data>{findObjectById($currentProject?.contactPoint[0].id)?.givenName?.split(";").join(" ")} {findObjectById($currentProject?.contactPoint[0].id)?.familyName}</div>
  {/if}
  {#if Array.isArray(findObjectById($currentProject?.contactPoint[0].id)?.memberOf)}
    {#each findObjectById($currentProject?.contactPoint[0].id)?.memberOf as o}
      <span>{findObjectById(o.id).name[0]}</span>
    {/each}
  {/if}
  {#if findObjectById($currentProject?.contactPoint[0].id)?.email}
    {#if Array.isArray(findObjectById($currentProject?.contactPoint[0].id)?.email)}
      <a class="data email" href="mailto:{findObjectById($currentProject?.contactPoint[0].id)?.email[0]}">{findObjectById($currentProject?.contactPoint[0].id)?.email[0]}</a>
    {:else}
      <a class="data email" href="mailto:{findObjectById($currentProject?.contactPoint[0].id)?.email}">{findObjectById($currentProject?.contactPoint[0].id)?.email}</a>
    {/if}
  {/if}
{/if}

<div class=label>Project Website</div>
{#if Array.isArray($currentProject?.url)}
  {#each $currentProject?.url as url}
    <a class="data external-link" href={url.url} target=_>{truncateString(url.name)}</a>
  {/each}
{/if}

{#if $currentProject}
  <div class=label>Keywords</div>
  <span class="keyword">{$currentProject?.keywords.join(", ")}</span>
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
