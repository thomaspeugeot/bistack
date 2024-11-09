import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OtherstackComponent } from './otherstack.component';

describe('OtherstackComponent', () => {
  let component: OtherstackComponent;
  let fixture: ComponentFixture<OtherstackComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OtherstackComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(OtherstackComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
