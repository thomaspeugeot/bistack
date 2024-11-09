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
  ) { }

  ngOnInit(): void {

    this.frontRepoService.connectToWebSocket(this.StackInstanceName).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log(this.StackInstanceName, "received web socket update", this.frontRepo.array_Foos[0].Name)

      }
    )

    console.log("ngOnInit", this.StackInstanceName)
  }

}
