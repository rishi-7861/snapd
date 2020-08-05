// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package boot

import (
	"github.com/snapcore/snapd/asserts"
	"github.com/snapcore/snapd/gadget"
)

func InitialSealObserver(model *asserts.Model) gadget.ContentWriteObserver {
	if model.Grade() == asserts.ModelGradeUnset {
		// no need to observe updates when assets are not managed
		return nil
	}

	return &initialSealObserver{
		model: model,
	}
}

type initialSealObserver struct {
	model *asserts.Model
}

// Implements gadget.ContentWriteObserver.
func (o *initialSealObserver) Observe(op gadget.ObserveAction, affectedStruct *gadget.LaidOutStructure, root, realSource, relativeTarget string) (gadget.ObserveResult, error) {
	// TODO:UC20:
	// steps on write action:
	// - copy new asset to assets cache
	// - update modeeenv
	// steps on rollback action:
	// - drop file from cache if no longer referenced
	// - update modeenv
	return gadget.ObserveResultNoted, nil
}

// Implements gadget.ContentWriteObserver.
func (o *initialSealObserver) Apply() error {
	// TODO:UC20: steps:
	// - reseal
	return nil
}
