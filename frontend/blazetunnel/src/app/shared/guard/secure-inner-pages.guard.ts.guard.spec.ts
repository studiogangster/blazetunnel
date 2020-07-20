import { TestBed } from '@angular/core/testing';

import { SecureInnerPages.Guard.TsGuard } from './secure-inner-pages.guard.ts.guard';

describe('SecureInnerPages.Guard.TsGuard', () => {
  let guard: SecureInnerPages.Guard.TsGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(SecureInnerPages.Guard.TsGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
