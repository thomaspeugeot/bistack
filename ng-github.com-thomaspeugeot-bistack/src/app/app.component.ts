import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as bistack from '../../projects/bistack/src/public-api'
import * as otherstack from '../../../otherstack/ng-github.com-thomaspeugeot-bistack-otherstack/projects/otherstack/src/public-api'

import { BistackspecificComponent } from '../../projects/bistackspecific/src/public-api'
import { OtherstackspecificComponent } from '../../../otherstack/ng-github.com-thomaspeugeot-bistack-otherstack/projects/otherstackspecific/src/public-api'

import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';


@Component({
    selector: 'app-root',
    imports: [
        CommonModule,
        FormsModule,
        MatRadioModule,
        MatButtonModule,
        MatIconModule,
        AngularSplitModule,
        BistackspecificComponent,
        OtherstackspecificComponent,
    ],
    templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {

  StacksNames = bistack.StacksNames
  OtherStacksNames = otherstack.StacksNames

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "bistack"
  StackType = bistack.StackType

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
