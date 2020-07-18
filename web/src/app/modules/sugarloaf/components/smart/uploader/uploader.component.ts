// Copyright (c) 2020 the Octant contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

import { Component, OnDestroy, OnInit, AfterViewInit } from '@angular/core';
import { WebsocketService } from '../../../../shared/services/websocket/websocket.service';
import { Subscription } from 'rxjs';
import { ContentService } from '../../../../shared/services/content/content.service';

@Component({
  selector: 'app-uploader',
  templateUrl: './uploader.component.html',
  styleUrls: ['./uploader.component.scss'],
})
export class UploaderComponent implements OnInit, OnDestroy, AfterViewInit {
  inputValue: string;
  showModal: boolean;
  contentCount = 0;

  private contentSubscription: Subscription;

  constructor(
    private websocketService: WebsocketService,
    private contentService: ContentService
  ) {}

  ngOnInit(): void {
    this.contentSubscription = this.contentService.current.subscribe(
      contentResponse => {
        this.contentCount++;
      }
    );

    this.websocketService.registerHandler('event.octant.dev/loading', () => {
      this.showModal = this.contentCount === 1;
    });
    this.websocketService.registerHandler('event.octant.dev/refresh', () => {
      setTimeout(window.location.reload.bind(window.location), 1000);
    });
  }

  ngAfterViewInit(): void {
    if (this.contentCount === 1) {
      this.websocketService.sendMessage('action.octant.dev/loading', {
        loading: true,
      });
    }
  }

  ngOnDestroy(): void {
    if (this.contentSubscription) {
      this.contentSubscription.unsubscribe();
    }
  }

  upload() {
    this.websocketService.sendMessage('action.octant.dev/uploadKubeConfig', {
      kubeConfig: window.btoa(this.inputValue),
    });
  }

  updateInput(event: HTMLInputElement) {
    this.inputValue = String(event);
  }

  hasInput(): boolean {
    return !this.inputValue || this.inputValue.length === 0;
  }
}
