<script lang="ts">
  interface Status {

  }
  import { statusFilter, query, getProjectsMetadata } from '../store';
  // import type { Category } from '../interfaces';
//   import { getProjectsMetadata } from '../store';

//   let categories = [
//     { id: 1, isOpen: false, name: 'Discipline', sub: ['Agriculture', 'Antropology', 'Geography', 'History'] },
//     { id: 2, isOpen: false, name: 'Type of data', sub: ['First', 'Second'] },
//     { id: 3, isOpen: false, name: 'Temporal coverage', sub: [] },
//     { id: 4, isOpen: false, name: 'Spatial coverage', sub: [] },
//     { id: 5, isOpen: false, name: 'Language', sub: [] },
//     { id: 6, isOpen: false, name: 'Keywords', sub: [] },
//     { id: 7, isOpen: false, name: 'Person', sub: [] },
//     { id: 8, isOpen: false, name: 'Organization', sub: ['Last', 'Not least'] },
//   ];

interface Category {
    id: number;
    isOpen: boolean;
    name: string;
    sub: SubCategory[];
  }
  interface SubCategory {
    isSelected: boolean;
    name: string;
  }

  let showFilters = false;

  // let categories = [
  //   {id: 1, isOpen: true, name: 'Status', sub: [
  //     {isSelected: $statusFilter.showInPlanning, name: 'In Planning'},
  //     {isSelected: $statusFilter.showOngoing, name: 'Ongoing'},
  //     {isSelected: $statusFilter.showFinished, name: 'Finished'}
  //   ]}
  // ]

  // const toggleCetegory = (cat: Category) => (event: MouseEvent) => {
  //   let bool = cat.isOpen;
  //   categories[cat.id - 1].isOpen = !bool;
  // };

  const toggleShowInPlanning = () => {
    $statusFilter.showInPlanning = !$statusFilter.showInPlanning;
    getProjectsMetadata(1, $query)
  };

  const toggleShowOngoing = () => {
    $statusFilter.showOngoing = !$statusFilter.showOngoing;
    getProjectsMetadata(1, $query)
  };

  const toggleShowFinished = () => {
    $statusFilter.showFinished = !$statusFilter.showFinished;
    getProjectsMetadata(1, $query)
  };
  
</script>

  <button on:click={() => {showFilters = !showFilters}}>
    Project Status
  </button>
  <div class={showFilters ? 'visible' : 'hidden'}>
    <label class=subcategory>
      <input on:click={toggleShowInPlanning} value={0} type=checkbox name=subcategory checked={$statusFilter.showInPlanning} />
      In Planning
    </label>
    <label class=subcategory>
      <input on:click={toggleShowOngoing} value={0} type=checkbox name=subcategory checked={$statusFilter.showOngoing} />
      Ongoing
    </label>
    <label class=subcategory>
      <input on:click={toggleShowFinished} value={0} type=checkbox name=subcategory checked={$statusFilter.showFinished} />
      Finished
    </label>
  </div>

<!-- {#each categories as category }
  <button on:click={toggleCetegory(category)} disabled={!category.sub.length}>
    {category.name}
  </button>
  {#if category.sub && category.sub.length}
    <div class={category.isOpen ? 'visible' : 'hidden'}>
      {#each category.sub as sub, n}
        <label class=subcategory>
          <input on:click={() => getProjectsMetadata(1, sub)} value={n} type=checkbox name=subcategory />
          {sub}
        </label>
      {/each}
    </div>
  {/if}
{/each} -->

<style>
  button {
    width: 100%;
    margin: 2px 0;
    padding: 10px;
    border: 1px solid #aaa;
    border-radius: 3px;
    text-align: left;
  }
  .subcategory {
    display: flex;
    align-items: center;
    cursor: pointer;
    margin: 5px 5px 5px 5px;
    padding: 5px;
    border-radius: 3px;
    background-color: #f2f2f2;
    font-size: 0.8em;
  }
  input[type=checkbox] {
    margin: 5px 10px;
    display: flex;
  }
  @media screen and (min-width: 992px) {
    button {
      padding: 5px 20px;
      margin: 5px;
      min-width: 200px;
    }
    .visible {
      display: block;
      width: 220px;
    }
  }
</style>
