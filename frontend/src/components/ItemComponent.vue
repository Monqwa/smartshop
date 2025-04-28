<template>
  <li class="item-component" :class="{ bought: item.is_buy }">
    <div class="item-checkbox-container" @click="toggleBoughtStatus">
      <input 
        type="checkbox" 
        :checked="item.is_buy" 
        class="item-checkbox"
        title="Toggle bought status"
      />
      <span class="checkbox-custom"></span>
    </div>
    
    <div class="item-content" @dblclick="startEditingName" title="Double-click to edit item name">
      <span v-if="!isEditingName" class="item-label">
        {{ item.name }}
      </span>
      <input 
        v-else 
        type="text" 
        v-model="editableName" 
        @blur="finishEditingName" 
        @keyup.enter="finishEditingName"
        @keyup.esc="cancelEditingName"
        ref="nameInput"
        class="item-name-input"
      />
    </div>

    <div class="item-actions">
      <button 
        v-if="!item.is_buy" 
        @click="startEditingName" 
        class="edit-btn" 
        title="Edit item name"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil" viewBox="0 0 16 16">
          <path d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5 13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"/>
        </svg>
      </button>
      <button @click="deleteItem" class="delete-btn" title="Delete item">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">
          <path d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5ZM11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H2.506a.58.58 0 0 0-.01 0H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1h-.995a.59.59 0 0 0-.01 0H11Zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5h9.916Zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47ZM8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z"/>
        </svg>
      </button>
    </div>
  </li>
</template>

<script>
import { ref, nextTick, watch } from 'vue';

export default {
  props: {
    item: {
      type: Object,
      required: true,
    }
  },
  emits: ['toggle-bought', 'delete-item', 'update-item-name'], 
  setup(props, { emit }) {
    const isEditingName = ref(false);
    const editableName = ref(props.item.name);
    const nameInput = ref(null);

    // Add watcher to update editableName when item.name changes
    watch(() => props.item.name, (newName) => {
      editableName.value = newName;
    });

    function toggleBoughtStatus() {
      emit('toggle-bought', props.item.id, !props.item.is_buy);
    }

    function deleteItem() {
      emit('delete-item', props.item.id);
    }

    async function startEditingName() {
      if (props.item.is_buy) return; // Don't allow editing bought items
      isEditingName.value = true;
      editableName.value = props.item.name; // Reset editable name to current prop value
      // Wait for the input to be rendered, then focus it
      await nextTick(); 
      nameInput.value?.focus();
    }

    function finishEditingName() {
      if (!isEditingName.value) return;
      const newName = editableName.value.trim();
      // Only emit update if name changed and is not empty
      if (newName && newName !== props.item.name) {
        emit('update-item-name', props.item.id, newName);
      }
      isEditingName.value = false;
    }

    function cancelEditingName() {
       isEditingName.value = false;
    }

    return {
      isEditingName,
      editableName,
      nameInput,
      toggleBoughtStatus,
      deleteItem,
      startEditingName,
      finishEditingName,
      cancelEditingName
    };
  }
};
</script>

<style scoped>
.item-component {
  display: flex;
  align-items: center;
  padding: var(--space-md);
  border: var(--border-width-thin) solid var(--item-border);
  border-radius: var(--border-radius-md);
  list-style: none;
  transition: all var(--transition-normal);
  position: relative;
  margin-bottom: var(--space-sm);
  background-color: var(--item-background);
  box-shadow: var(--shadow-sm);
}

.item-component:hover {
  background-color: var(--item-hover-background);
  transform: translateY(-1px);
  box-shadow: var(--shadow-md);
}

.item-checkbox-container {
  position: relative;
  margin-right: var(--space-md);
  cursor: pointer;
  width: 22px;
  height: 22px;
  flex-shrink: 0;
}

.item-checkbox {
  position: absolute;
  opacity: 0;
  cursor: pointer;
  height: 0;
  width: 0;
}

.checkbox-custom {
  position: absolute;
  top: 0;
  left: 0;
  height: 22px;
  width: 22px;
  border-radius: var(--border-radius-sm);
  background-color: var(--checkbox-background);
  border: var(--border-width-thin) solid var(--checkbox-border);
  transition: all var(--transition-normal);
}

.item-checkbox:checked ~ .checkbox-custom {
  background-color: var(--checkbox-checked-background);
  border-color: var(--checkbox-checked-border);
}

.checkbox-custom:after {
  content: "";
  position: absolute;
  display: none;
  left: 8px;
  top: 4px;
  width: 5px;
  height: 10px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.item-checkbox:checked ~ .checkbox-custom:after {
  display: block;
}

.item-checkbox-container:hover .checkbox-custom {
  border-color: var(--checkbox-active);
}

.item-content {
  flex-grow: 1;
  min-width: 0;
  cursor: text;
  position: relative;
  padding: var(--space-xs) 0;
}

.item-label {
  display: block;
  padding: var(--space-xs) 0;
  color: var(--text-color);
  font-size: var(--font-size-md);
  line-height: 1.5;
  transition: all var(--transition-normal);
  word-break: break-word;
  font-weight: var(--font-weight-medium);
  letter-spacing: 0.01em;
}

.item-component:not(.bought):hover .item-label {
  color: var(--color-gray-800);
}

.item-name-input {
  font-size: var(--font-size-md);
  line-height: 1.5;
  padding: var(--space-xs) var(--space-sm);
  margin: 0;
  border: var(--border-width-thin) solid var(--item-focus-border);
  border-radius: var(--border-radius-sm);
  width: 100%;
  box-sizing: border-box;
  font-family: var(--font-family);
  background-color: var(--color-white);
  font-weight: var(--font-weight-medium);
  letter-spacing: 0.01em;
}

.item-name-input:focus {
  outline: none;
  box-shadow: var(--shadow-md);
  border-color: var(--primary);
}

.item-component.bought {
  opacity: 0.8;
  background-color: var(--color-gray-50);
}

.item-component.bought .item-label {
  text-decoration: line-through;
  color: var(--item-bought-color);
  font-style: italic;
}

.item-component.bought .item-content {
  cursor: default;
}

.item-actions {
  display: flex;
  gap: var(--space-xs);
  margin-left: var(--space-md);
}

.edit-btn,
.delete-btn {
  background: none;
  border: none;
  color: var(--color-gray-400);
  font-size: var(--font-size-md);
  cursor: pointer;
  padding: var(--space-sm);
  line-height: 1;
  opacity: 0.4;
  transition: all var(--transition-normal);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.item-component:hover .edit-btn,
.item-component:hover .delete-btn {
  opacity: 0.7;
}

.edit-btn:hover {
  opacity: 1;
  color: var(--primary);
  background-color: var(--color-gray-100);
  transform: scale(1.1);
}

.delete-btn:hover {
  opacity: 1;
  color: var(--danger);
  background-color: var(--delete-button-hover);
  transform: scale(1.1);
}

.edit-btn:focus,
.delete-btn:focus {
  outline: none;
  box-shadow: var(--shadow-sm);
}

@media (max-width: 600px) {
  .item-component {
    padding: var(--space-sm);
  }
  
  .item-label {
    font-size: var(--font-size-sm);
  }
  
  .edit-btn,
  .delete-btn {
    opacity: 0.5;
  }
  
  .item-actions {
    margin-left: var(--space-sm);
  }
}
</style> 