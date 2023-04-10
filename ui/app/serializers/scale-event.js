/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import ApplicationSerializer from './application';
import classic from 'ember-classic-decorator';

@classic
export default class ScaleEventSerializer extends ApplicationSerializer {
  separateNanos = ['Time'];
  objectNullOverrides = ['Meta'];
}
