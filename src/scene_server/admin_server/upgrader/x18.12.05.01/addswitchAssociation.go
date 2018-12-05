/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package x18_12_05_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/metadata"
	"configcenter/src/common/util"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addswitchAssociation(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	switchAsst := metadata.Association{
		OwnerID:         conf.OwnerID,
		AsstKindID:      "connect",
		ObjectID:        "bk_switch",
		AsstObjID:       "host",
		AssociationName: "bk_switch_host",
		Mapping:         metadata.OneToOneMapping,
		OnDelete:        metadata.NoAction,
		IsPre:           util.Pfalse(),
	}

	_, _, err := upgrader.Upsert(ctx, db, common.BKTableNameObjAsst, switchAsst, "id", []string{"bk_obj_id", "bk_asst_obj_id"}, []string{"id"})
	if err != nil {
		return err
	}

	return nil
}
