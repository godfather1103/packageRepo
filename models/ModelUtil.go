package models

func GetWebTitle(pathFileId int64, idType int64) string {
	if idType == 0 || idType == 1 {
		pathInfo, err := FindPathInfoById(pathFileId)
		if err == nil {
			if pathInfo.ParentPathId == 0 {
				return pathInfo.PathName
			} else {
				return GetWebTitle(pathInfo.ParentPathId, 0) + "/" + pathInfo.PathName
			}
		}
	} else if idType == 2 {
		versioInfo, err := FindVersionInfoById(pathFileId)
		if err == nil {
			return GetWebTitle(versioInfo.ArtifactId, 1) + "/" + versioInfo.Version
		}
	}
	return "首页"
}

func GetBreadcrumb(pathFileId int64, idType int64) []*PathAndVersion {
	if idType == 0 || idType == 1 {
		pathInfo, err := FindPathInfoById(pathFileId)
		if err == nil {
			var nowIdType int64
			if pathInfo.PathType == 2 {
				nowIdType = 1
			} else {
				nowIdType = 0
			}
			if pathInfo.ParentPathId == 0 {
				var result = make([]*PathAndVersion, 1)
				var pathInfoMap = &PathAndVersion{
					PATHFILEID: pathInfo.Id,
					IDTYPE:     nowIdType,
					PATHNAME:   pathInfo.PathName,
				}
				result[0] = pathInfoMap
				return result
			} else {
				result1 := GetBreadcrumb(pathInfo.ParentPathId, 0)
				result := make([]*PathAndVersion, len(result1)+1)
				var pathInfoMap = &PathAndVersion{
					PATHFILEID: pathInfo.Id,
					IDTYPE:     nowIdType,
					PATHNAME:   pathInfo.PathName,
				}
				for i := 0; i < len(result1); i++ {
					result[i] = result1[i]
				}
				result[len(result1)] = pathInfoMap
				return result
			}
		}
	} else if idType == 2 {
		versioInfo, err := FindVersionInfoById(pathFileId)
		if err == nil {
			result1 := GetBreadcrumb(versioInfo.ArtifactId, 1)
			result := make([]*PathAndVersion, len(result1)+1)
			var pathInfoMap = &PathAndVersion{
				PATHFILEID: versioInfo.Id,
				IDTYPE:     2,
				PATHNAME:   versioInfo.Version,
			}
			for i := 0; i < len(result1); i++ {
				result[i] = result1[i]
			}
			result[len(result1)] = pathInfoMap
			return result
		}
	}
	return make([]*PathAndVersion, 0)
}
