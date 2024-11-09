import { Component, Input, OnInit } from '@angular/core';

import * as bistack from '../../../bistack/src/public-api'
import { MatTableModule } from '@angular/material/table';

@Component({
  selector: 'lib-bistackspecific',
  standalone: true,
  imports: [
    MatTableModule
  ],
  templateUrl: './bistackspecific.component.html',
  styles: ``
})
export class BistackspecificComponent implements OnInit {

  @Input() StackInstanceName: string = ""

  public frontRepo = new (bistack.FrontRepo)

  constructor(
    private frontRepoService: bistack.FrontRepoService,

    private fooService: bistack.FooService,

  ) { }

  ngOnInit(): void {

    this.fooService.getFoos(this.StackInstanceName, this.frontRepo).subscribe(
      foos => {
        this.frontRepo.array_Foos = foos
      }
    )

    this.frontRepoService.connectToWebSocket(this.StackInstanceName).subscribe(
      gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo
      }
    )

    console.log("ngOnInit")
  }

}
