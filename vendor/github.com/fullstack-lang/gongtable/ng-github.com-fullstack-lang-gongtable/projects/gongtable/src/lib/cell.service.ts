// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { CellAPI } from './cell-api'
import { Cell, CopyCellToCellAPI } from './cell'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { CellStringAPI } from './cellstring-api'
import { CellFloat64API } from './cellfloat64-api'
import { CellIntAPI } from './cellint-api'
import { CellBooleanAPI } from './cellboolean-api'
import { CellIconAPI } from './cellicon-api'

@Injectable({
  providedIn: 'root'
})
export class CellService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CellServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private cellsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.cellsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/cells';
  }

  /** GET cells from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI[]> {
    return this.getCells(GONG__StackPath, frontRepo)
  }
  getCells(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CellAPI[]>(this.cellsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CellAPI[]>('getCells', []))
      );
  }

  /** GET cell by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI> {
    return this.getCell(id, GONG__StackPath, frontRepo)
  }
  getCell(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.cellsUrl}/${id}`;
    return this.http.get<CellAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched cell id=${id}`)),
      catchError(this.handleError<CellAPI>(`getCell id=${id}`))
    );
  }

  // postFront copy cell to a version with encoded pointers and post to the back
  postFront(cell: Cell, GONG__StackPath: string): Observable<CellAPI> {
    let cellAPI = new CellAPI
    CopyCellToCellAPI(cell, cellAPI)
    const id = typeof cellAPI === 'number' ? cellAPI : cellAPI.ID
    const url = `${this.cellsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellAPI>(url, cellAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CellAPI>('postCell'))
    );
  }
  
  /** POST: add a new cell to the server */
  post(celldb: CellAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI> {
    return this.postCell(celldb, GONG__StackPath, frontRepo)
  }
  postCell(celldb: CellAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellAPI>(this.cellsUrl, celldb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted celldb id=${celldb.ID}`)
      }),
      catchError(this.handleError<CellAPI>('postCell'))
    );
  }

  /** DELETE: delete the celldb from the server */
  delete(celldb: CellAPI | number, GONG__StackPath: string): Observable<CellAPI> {
    return this.deleteCell(celldb, GONG__StackPath)
  }
  deleteCell(celldb: CellAPI | number, GONG__StackPath: string): Observable<CellAPI> {
    const id = typeof celldb === 'number' ? celldb : celldb.ID;
    const url = `${this.cellsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CellAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted celldb id=${id}`)),
      catchError(this.handleError<CellAPI>('deleteCell'))
    );
  }

  // updateFront copy cell to a version with encoded pointers and update to the back
  updateFront(cell: Cell, GONG__StackPath: string): Observable<CellAPI> {
    let cellAPI = new CellAPI
    CopyCellToCellAPI(cell, cellAPI)
    const id = typeof cellAPI === 'number' ? cellAPI : cellAPI.ID
    const url = `${this.cellsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<CellAPI>(url, cellAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CellAPI>('updateCell'))
    );
  }

  /** PUT: update the celldb on the server */
  update(celldb: CellAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI> {
    return this.updateCell(celldb, GONG__StackPath, frontRepo)
  }
  updateCell(celldb: CellAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellAPI> {
    const id = typeof celldb === 'number' ? celldb : celldb.ID;
    const url = `${this.cellsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CellAPI>(url, celldb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated celldb id=${celldb.ID}`)
      }),
      catchError(this.handleError<CellAPI>('updateCell'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CellService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CellService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}