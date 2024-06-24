package queries

func RetrieveChildItems() string {
	return `
	SELECT  PA.name 'pa_name',
				CASE WHEN PA.objtype_id = 1562 THEN 'Location'
					WHEN PA.objtype_id = 4 THEN 'Server'
					WHEN PA.objtype_id = 1502 THEN 'Chassis'
					WHEN PA.objtype_id = 1561 THEN 'Row'
					WHEN PA.objtype_id = 1560 THEN 'Rack'
					ELSE 'Other'
					END AS 'pa_type',
 		A.child_entity_type 'ch_entity', CH.name 'ch_name',
 				CASE WHEN CH.objtype_id = 4 THEN 'Server'
 					WHEN CH.objtype_id = 1504 THEN 'VM'
 					WHEN CH.objtype_id = 1561 THEN 'Row'
 					WHEN CH.objtype_id = 1502 THEN 'Chassis'
 					WHEN CH.objtype_id = 1562 THEN 'Location'
 					WHEN CH.objtype_id = 1560 THEN 'Rack'
 					ELSE 'Other'
 					END AS 'ch_type'
FROM racktables.EntityLink A
LEFT JOIN racktables.Object CH ON A.child_entity_id = CH.id
LEFT JOIN racktables.Object PA ON A.parent_entity_id = PA.id
ORDER BY CH.objtype_id ASC, CH.name ASC
LIMIT 1000
;`
}
