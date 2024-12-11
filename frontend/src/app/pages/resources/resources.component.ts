import { CommonModule } from '@angular/common';
import { Component, signal } from '@angular/core';
import { Task } from './../../models/task.model'
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms'

@Component({
  selector: 'app-resources',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './resources.component.html',
  styleUrl: './resources.component.css'
})
export class ResourcesComponent {
  name = signal("brayan")

  colorCtrl = new FormControl()

  widthCtrl = new FormControl(50, {
    nonNullable: true,

  })

  nameCtrl = new FormControl('brayan', {
    nonNullable: true,
    validators: [
      Validators.required,
      Validators.minLength(3)
    ]
  })

  constructor() {
    this.colorCtrl.valueChanges.subscribe(value => {
      console.log(value)
    })
  } 

  person = signal({
    name: 'e',
    age: 18,
  })

  otherTasks = signal<Task[]>([
    {
      id: Date.now(),
      title: 'pepe',
      completed: false,
    },
    {
      id: Date.now(),
      title: 'pepe2',
      completed: false,
    }
  ])

  newTaskControl = new FormControl('', {
    nonNullable: true,
    validators: [
      Validators.required,
    ]
  })

  changeHandler() {
    if (this.newTaskControl.valid){
      const value = this.newTaskControl.value.trim();

      if (value !== '') {
        this.addTask(value)
        this.newTaskControl.setValue('')
      }
    }
  }

  addTask(title: string) {
    const newTask = {
      id: Date.now(),
      title: title,
      completed: false,
    };

    this.otherTasks.update((tasks) => [...tasks, newTask])
  }

  destroyHandler(index: number) {
    this.otherTasks.update((tasks) => tasks.filter((task, pos) => pos !== index))
  }

  updateCompletedHandler(index: number) {
    this.otherTasks.update((tasks) => {
      tasks[index].completed = !tasks[index].completed

      return tasks
    })
  }

  changeName(event: Event) {
    const input = event.target as HTMLInputElement
    const value = input.value

    this.person.update(prevState => {
      return {
        ...prevState,
        name: value,
      }
    })
  }
}
