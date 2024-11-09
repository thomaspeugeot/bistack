import { TestBed } from '@angular/core/testing';

import { OtherstackspecificService } from './otherstackspecific.service';

describe('OtherstackspecificService', () => {
  let service: OtherstackspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OtherstackspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
