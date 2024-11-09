import { TestBed } from '@angular/core/testing';

import { OtherstackService } from './otherstack.service';

describe('OtherstackService', () => {
  let service: OtherstackService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OtherstackService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
