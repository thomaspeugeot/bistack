import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OtherstackspecificComponent } from './otherstackspecific.component';

describe('OtherstackspecificComponent', () => {
  let component: OtherstackspecificComponent;
  let fixture: ComponentFixture<OtherstackspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OtherstackspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(OtherstackspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
