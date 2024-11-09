import { Component } from '@angular/core';

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
export class BistackspecificComponent {

  StacksNames = bistack.StacksNames

  public frontRepo = new (bistack.FrontRepo)

  constructor(
    private frontRepoService: bistack.FrontRepoService,

    private fooService: bistack.FooService,

  ) { }

  ngOnInit(): void {

    this.fooService.getFoos(this.StacksNames.Bistack, this.frontRepo).subscribe(
      foos => {
        this.frontRepo.array_Foos = foos
      }
    )

    this.frontRepoService.connectToWebSocket(this.StacksNames.Bistack).subscribe(
      gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo
      }
    )

    console.log("ngOnInit")
  }

}
