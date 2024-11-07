import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BistackspecificComponent } from './bistackspecific.component';

describe('BistackspecificComponent', () => {
  let component: BistackspecificComponent;
  let fixture: ComponentFixture<BistackspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BistackspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(BistackspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
