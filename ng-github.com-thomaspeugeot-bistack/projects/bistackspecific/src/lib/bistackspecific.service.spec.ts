import { TestBed } from '@angular/core/testing';

import { BistackspecificService } from './bistackspecific.service';

describe('BistackspecificService', () => {
  let service: BistackspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BistackspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
