

import { Component, Input, OnInit } from '@angular/core';

import * as otherstack from '../../../otherstack/src/public-api'
import { MatTableModule } from '@angular/material/table';

@Component({
  selector: 'lib-otherstackspecific',
  standalone: true,
  imports: [
    MatTableModule
  ],
  templateUrl: "otherstackspecific.component.html",
  styles: ``
})
export class OtherstackspecificComponent implements OnInit {

  @Input() StackInstanceName: string = ""

  public frontRepo = new (otherstack.FrontRepo)

  StackType = otherstack.StackType

  constructor(
    private frontRepoService: otherstack.FrontRepoService,
  ) { }

  ngOnInit(): void {

    this.frontRepoService.connectToWebSocket(this.StackInstanceName).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        // console.log(this.StackInstanceName, "received web socket update", this.frontRepo.array_Bars[0].Name)

      }
    )

    console.log("ngOnInit", this.StackInstanceName)
  }

}
