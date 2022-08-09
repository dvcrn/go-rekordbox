CREATE TABLE `agentRegistry` (
  `registry_id` VARCHAR(255) PRIMARY KEY, 
  `id_1` VARCHAR(255) DEFAULT NULL, 
  `id_2` VARCHAR(255) DEFAULT NULL, 
  `int_1` BIGINT DEFAULT NULL, 
  `int_2` BIGINT DEFAULT NULL, 
  `str_1` VARCHAR(255) DEFAULT NULL, 
  `str_2` VARCHAR(255) DEFAULT NULL, 
  `date_1` DATETIME DEFAULT NULL, 
  `date_2` DATETIME DEFAULT NULL, 
  `text_1` TEXT DEFAULT NULL, 
  `text_2` TEXT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `cloudAgentRegistry` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `int_1` BIGINT DEFAULT NULL, 
  `int_2` BIGINT DEFAULT NULL, 
  `str_1` VARCHAR(255) DEFAULT NULL, 
  `str_2` VARCHAR(255) DEFAULT NULL, 
  `date_1` DATETIME DEFAULT NULL, 
  `date_2` DATETIME DEFAULT NULL, 
  `text_1` TEXT DEFAULT NULL, 
  `text_2` TEXT DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `contentActiveCensor` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `ActiveCensors` TEXT DEFAULT NULL, 
  `rb_activecensor_count` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `contentCue` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `Cues` TEXT DEFAULT NULL, 
  `rb_cue_count` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `contentFile` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `Path` VARCHAR(255) DEFAULT NULL, 
  `Hash` VARCHAR(255) DEFAULT NULL, 
  `Size` INTEGER DEFAULT NULL, 
  `rb_local_path` VARCHAR(255) DEFAULT NULL, 
  `rb_insync_hash` VARCHAR(255) DEFAULT NULL, 
  `rb_insync_local_usn` BIGINT DEFAULT NULL, 
  `rb_file_hash_dirty` INTEGER DEFAULT 0, 
  `rb_local_file_status` INTEGER DEFAULT 0, 
  `rb_in_progress` TINYINT(1) DEFAULT 0, 
  `rb_process_type` INTEGER DEFAULT 0, 
  `rb_temp_path` VARCHAR(255) DEFAULT NULL, 
  `rb_priority` INTEGER DEFAULT 50, 
  `rb_file_size_dirty` INTEGER DEFAULT 0, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdActiveCensor` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `InMsec` INTEGER DEFAULT NULL, 
  `OutMsec` INTEGER DEFAULT NULL, 
  `Info` INTEGER DEFAULT NULL, 
  `ParameterList` TEXT DEFAULT NULL, 
  `ContentUUID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdAlbum` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `AlbumArtistID` VARCHAR(255) DEFAULT NULL, 
  `ImagePath` VARCHAR(255) DEFAULT NULL, 
  `Compilation` INTEGER DEFAULT NULL, 
  `SearchStr` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdArtist` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `SearchStr` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdCategory` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `MenuItemID` VARCHAR(255) DEFAULT NULL, 
  `Seq` INTEGER DEFAULT NULL, 
  `Disable` INTEGER DEFAULT NULL, 
  `InfoOrder` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdCloudProperty` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Reserved1` TEXT DEFAULT NULL, 
  `Reserved2` TEXT DEFAULT NULL, 
  `Reserved3` TEXT DEFAULT NULL, 
  `Reserved4` TEXT DEFAULT NULL, 
  `Reserved5` TEXT DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdColor` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ColorCode` INTEGER DEFAULT NULL, 
  `SortKey` INTEGER DEFAULT NULL, 
  `Commnt` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdContent` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `FolderPath` VARCHAR(255) DEFAULT NULL, 
  `FileNameL` VARCHAR(255) DEFAULT NULL, 
  `FileNameS` VARCHAR(255) DEFAULT NULL, 
  `Title` VARCHAR(255) DEFAULT NULL, 
  `ArtistID` VARCHAR(255) DEFAULT NULL, 
  `AlbumID` VARCHAR(255) DEFAULT NULL, 
  `GenreID` VARCHAR(255) DEFAULT NULL, 
  `BPM` INTEGER DEFAULT NULL, 
  `Length` INTEGER DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `BitRate` INTEGER DEFAULT NULL, 
  `BitDepth` INTEGER DEFAULT NULL, 
  `Commnt` TEXT DEFAULT NULL, 
  `FileType` INTEGER DEFAULT NULL, 
  `Rating` INTEGER DEFAULT NULL, 
  `ReleaseYear` INTEGER DEFAULT NULL, 
  `RemixerID` VARCHAR(255) DEFAULT NULL, 
  `LabelID` VARCHAR(255) DEFAULT NULL, 
  `OrgArtistID` VARCHAR(255) DEFAULT NULL, 
  `KeyID` VARCHAR(255) DEFAULT NULL, 
  `StockDate` VARCHAR(255) DEFAULT NULL, 
  `ColorID` VARCHAR(255) DEFAULT NULL, 
  `DJPlayCount` INTEGER DEFAULT NULL, 
  `ImagePath` VARCHAR(255) DEFAULT NULL, 
  `MasterDBID` VARCHAR(255) DEFAULT NULL, 
  `MasterSongID` VARCHAR(255) DEFAULT NULL, 
  `AnalysisDataPath` VARCHAR(255) DEFAULT NULL, 
  `SearchStr` VARCHAR(255) DEFAULT NULL, 
  `FileSize` INTEGER DEFAULT NULL, 
  `DiscNo` INTEGER DEFAULT NULL, 
  `ComposerID` VARCHAR(255) DEFAULT NULL, 
  `Subtitle` VARCHAR(255) DEFAULT NULL, 
  `SampleRate` INTEGER DEFAULT NULL, 
  `DisableQuantize` INTEGER DEFAULT NULL, 
  `Analysed` INTEGER DEFAULT NULL, 
  `ReleaseDate` VARCHAR(255) DEFAULT NULL, 
  `DateCreated` VARCHAR(255) DEFAULT NULL, 
  `ContentLink` INTEGER DEFAULT NULL, 
  `Tag` VARCHAR(255) DEFAULT NULL, 
  `ModifiedByRBM` VARCHAR(255) DEFAULT NULL, 
  `HotCueAutoLoad` VARCHAR(255) DEFAULT NULL, 
  `DeliveryControl` VARCHAR(255) DEFAULT NULL, 
  `DeliveryComment` VARCHAR(255) DEFAULT NULL, 
  `CueUpdated` VARCHAR(255) DEFAULT NULL, 
  `AnalysisUpdated` VARCHAR(255) DEFAULT NULL, 
  `TrackInfoUpdated` VARCHAR(255) DEFAULT NULL, 
  `Lyricist` VARCHAR(255) DEFAULT NULL, 
  `ISRC` VARCHAR(255) DEFAULT NULL, 
  `SamplerTrackInfo` INTEGER DEFAULT NULL, 
  `SamplerPlayOffset` INTEGER DEFAULT NULL, 
  `SamplerGain` FLOAT DEFAULT NULL, 
  `VideoAssociate` VARCHAR(255) DEFAULT NULL, 
  `LyricStatus` INTEGER DEFAULT NULL, 
  `ServiceID` INTEGER DEFAULT NULL, 
  `OrgFolderPath` VARCHAR(255) DEFAULT NULL, 
  `Reserved1` TEXT DEFAULT NULL, 
  `Reserved2` TEXT DEFAULT NULL, 
  `Reserved3` TEXT DEFAULT NULL, 
  `Reserved4` TEXT DEFAULT NULL, 
  `ExtInfo` TEXT DEFAULT NULL, 
  `rb_file_id` VARCHAR(255) DEFAULT NULL, 
  `DeviceID` VARCHAR(255) DEFAULT NULL, 
  `rb_LocalFolderPath` VARCHAR(255) DEFAULT NULL, 
  `SrcID` VARCHAR(255) DEFAULT NULL, 
  `SrcTitle` VARCHAR(255) DEFAULT NULL, 
  `SrcArtistName` VARCHAR(255) DEFAULT NULL, 
  `SrcAlbumName` VARCHAR(255) DEFAULT NULL, 
  `SrcLength` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdCue` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `InMsec` INTEGER DEFAULT NULL, 
  `InFrame` INTEGER DEFAULT NULL, 
  `InMpegFrame` INTEGER DEFAULT NULL, 
  `InMpegAbs` INTEGER DEFAULT NULL, 
  `OutMsec` INTEGER DEFAULT NULL, 
  `OutFrame` INTEGER DEFAULT NULL, 
  `OutMpegFrame` INTEGER DEFAULT NULL, 
  `OutMpegAbs` INTEGER DEFAULT NULL, 
  `Kind` INTEGER DEFAULT NULL, 
  `Color` INTEGER DEFAULT NULL, 
  `ColorTableIndex` INTEGER DEFAULT NULL, 
  `ActiveLoop` INTEGER DEFAULT NULL, 
  `Comment` VARCHAR(255) DEFAULT NULL, 
  `BeatLoopSize` INTEGER DEFAULT NULL, 
  `CueMicrosec` INTEGER DEFAULT NULL, 
  `InPointSeekInfo` VARCHAR(255) DEFAULT NULL, 
  `OutPointSeekInfo` VARCHAR(255) DEFAULT NULL, 
  `ContentUUID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdDevice` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `MasterDBID` VARCHAR(255) DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdGenre` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdHistory` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Seq` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `Attribute` INTEGER DEFAULT NULL, 
  `ParentID` VARCHAR(255) DEFAULT NULL, 
  `DateCreated` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdHotCueBanklist` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Seq` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `ImagePath` VARCHAR(255) DEFAULT NULL, 
  `Attribute` INTEGER DEFAULT NULL, 
  `ParentID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdKey` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ScaleName` VARCHAR(255) DEFAULT NULL, 
  `Seq` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdLabel` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdMenuItems` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Class` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdMixerParam` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `GainHigh` INTEGER DEFAULT NULL, 
  `GainLow` INTEGER DEFAULT NULL, 
  `PeakHigh` INTEGER DEFAULT NULL, 
  `PeakLow` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdMyTag` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Seq` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `Attribute` INTEGER DEFAULT NULL, 
  `ParentID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdPlaylist` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Seq` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `ImagePath` VARCHAR(255) DEFAULT NULL, 
  `Attribute` INTEGER DEFAULT NULL, 
  `ParentID` VARCHAR(255) DEFAULT NULL, 
  `SmartList` TEXT DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdProperty` (
  `DBID` VARCHAR(255) PRIMARY KEY, 
  `DBVersion` VARCHAR(255) DEFAULT NULL, 
  `BaseDBDrive` VARCHAR(255) DEFAULT NULL, 
  `CurrentDBDrive` VARCHAR(255) DEFAULT NULL, 
  `DeviceID` VARCHAR(255) DEFAULT NULL, 
  `Reserved1` TEXT DEFAULT NULL, 
  `Reserved2` TEXT DEFAULT NULL, 
  `Reserved3` TEXT DEFAULT NULL, 
  `Reserved4` TEXT DEFAULT NULL, 
  `Reserved5` TEXT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdRecommendLike` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID1` VARCHAR(255) DEFAULT NULL, 
  `ContentID2` VARCHAR(255) DEFAULT NULL, 
  `LikeRate` INTEGER DEFAULT NULL, 
  `DataCreatedH` INTEGER DEFAULT NULL, 
  `DataCreatedL` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdRelatedTracks` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Seq` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `Attribute` INTEGER DEFAULT NULL, 
  `ParentID` VARCHAR(255) DEFAULT NULL, 
  `Criteria` TEXT DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSampler` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Seq` INTEGER DEFAULT NULL, 
  `Name` VARCHAR(255) DEFAULT NULL, 
  `Attribute` INTEGER DEFAULT NULL, 
  `ParentID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongHistory` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `HistoryID` VARCHAR(255) DEFAULT NULL, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongHotCueBanklist` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `HotCueBanklistID` VARCHAR(255) DEFAULT NULL, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `CueID` VARCHAR(255) DEFAULT NULL, 
  `InMsec` INTEGER DEFAULT NULL, 
  `InFrame` INTEGER DEFAULT NULL, 
  `InMpegFrame` INTEGER DEFAULT NULL, 
  `InMpegAbs` INTEGER DEFAULT NULL, 
  `OutMsec` INTEGER DEFAULT NULL, 
  `OutFrame` INTEGER DEFAULT NULL, 
  `OutMpegFrame` INTEGER DEFAULT NULL, 
  `OutMpegAbs` INTEGER DEFAULT NULL, 
  `Color` INTEGER DEFAULT NULL, 
  `ColorTableIndex` INTEGER DEFAULT NULL, 
  `ActiveLoop` INTEGER DEFAULT NULL, 
  `Comment` VARCHAR(255) DEFAULT NULL, 
  `BeatLoopSize` INTEGER DEFAULT NULL, 
  `CueMicrosec` INTEGER DEFAULT NULL, 
  `InPointSeekInfo` VARCHAR(255) DEFAULT NULL, 
  `OutPointSeekInfo` VARCHAR(255) DEFAULT NULL, 
  `HotCueBanklistUUID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongMyTag` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `MyTagID` VARCHAR(255) DEFAULT NULL, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongPlaylist` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `PlaylistID` VARCHAR(255) DEFAULT NULL, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongRelatedTracks` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `RelatedTracksID` VARCHAR(255) DEFAULT NULL, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongSampler` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `SamplerID` VARCHAR(255) DEFAULT NULL, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSongTagList` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `ContentID` VARCHAR(255) DEFAULT NULL, 
  `TrackNo` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `djmdSort` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `MenuItemID` VARCHAR(255) DEFAULT NULL, 
  `Seq` INTEGER DEFAULT NULL, 
  `Disable` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `hotCueBanklistCue` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `HotCueBanklistID` VARCHAR(255) DEFAULT NULL, 
  `Cues` TEXT DEFAULT NULL, 
  `rb_cue_count` INTEGER DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `settingFile` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `Path` VARCHAR(255) DEFAULT NULL, 
  `Hash` VARCHAR(255) DEFAULT NULL, 
  `Size` INTEGER DEFAULT NULL, 
  `rb_local_path` VARCHAR(255) DEFAULT NULL, 
  `rb_insync_hash` VARCHAR(255) DEFAULT NULL, 
  `rb_insync_local_usn` BIGINT DEFAULT NULL, 
  `rb_file_hash_dirty` INTEGER DEFAULT 0, 
  `rb_file_size_dirty` INTEGER DEFAULT 0, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `uuidIDMap` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `TableName` VARCHAR(255) DEFAULT NULL, 
  `TargetUUID` VARCHAR(255) DEFAULT NULL, 
  `CurrentID` VARCHAR(255) DEFAULT NULL, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE TABLE `imageFile` (
  `ID` VARCHAR(255) PRIMARY KEY, 
  `TableName` VARCHAR(255) DEFAULT NULL, 
  `TargetUUID` VARCHAR(255) DEFAULT NULL, 
  `TargetID` VARCHAR(255) DEFAULT NULL, 
  `Path` VARCHAR(255) DEFAULT NULL, 
  `Hash` VARCHAR(255) DEFAULT NULL, 
  `Size` INTEGER DEFAULT NULL, 
  `rb_local_path` VARCHAR(255) DEFAULT NULL, 
  `rb_insync_hash` VARCHAR(255) DEFAULT NULL, 
  `rb_insync_local_usn` BIGINT DEFAULT NULL, 
  `rb_file_hash_dirty` INTEGER DEFAULT 0, 
  `rb_local_file_status` INTEGER DEFAULT 0, 
  `rb_in_progress` TINYINT(1) DEFAULT 0, 
  `rb_process_type` INTEGER DEFAULT 0, 
  `rb_temp_path` VARCHAR(255) DEFAULT NULL, 
  `rb_priority` INTEGER DEFAULT 50, 
  `rb_file_size_dirty` INTEGER DEFAULT 0, 
  `UUID` VARCHAR(255) DEFAULT NULL, 
  `rb_data_status` INTEGER DEFAULT 0, 
  `rb_local_data_status` INTEGER DEFAULT 0, 
  `rb_local_deleted` TINYINT(1) DEFAULT 0, 
  `rb_local_synced` TINYINT(1) DEFAULT 0, 
  `usn` BIGINT DEFAULT NULL, 
  `rb_local_usn` BIGINT DEFAULT NULL, 
  `created_at` DATETIME NOT NULL, 
  `updated_at` DATETIME NOT NULL
);
CREATE INDEX `agent_registry_id_1_id_2` ON `agentRegistry` (`id_1`, `id_2`);
CREATE INDEX `cloud_agent_registry__u_u_i_d` ON `cloudAgentRegistry` (`UUID`);
CREATE INDEX `cloud_agent_registry_rb_data_status` ON `cloudAgentRegistry` (`rb_data_status`);
CREATE INDEX `cloud_agent_registry_rb_local_data_status` ON `cloudAgentRegistry` (`rb_local_data_status`);
CREATE INDEX `cloud_agent_registry_rb_local_deleted` ON `cloudAgentRegistry` (`rb_local_deleted`);
CREATE INDEX `cloud_agent_registry_rb_local_usn__i_d` ON `cloudAgentRegistry` (`rb_local_usn`, `ID`);
CREATE INDEX `content_active_censor__content_i_d` ON `contentActiveCensor` (`ContentID`);
CREATE INDEX `content_active_censor__u_u_i_d` ON `contentActiveCensor` (`UUID`);
CREATE INDEX `content_active_censor_rb_data_status` ON `contentActiveCensor` (`rb_data_status`);
CREATE INDEX `content_active_censor_rb_local_data_status` ON `contentActiveCensor` (`rb_local_data_status`);
CREATE INDEX `content_active_censor_rb_local_deleted` ON `contentActiveCensor` (`rb_local_deleted`);
CREATE INDEX `content_active_censor_rb_local_usn__i_d` ON `contentActiveCensor` (`rb_local_usn`, `ID`);
CREATE INDEX `content_cue_rb_cue_count` ON `contentCue` (`rb_cue_count`);
CREATE INDEX `content_cue__content_i_d` ON `contentCue` (`ContentID`);
CREATE INDEX `content_cue__u_u_i_d` ON `contentCue` (`UUID`);
CREATE INDEX `content_cue_rb_data_status` ON `contentCue` (`rb_data_status`);
CREATE INDEX `content_cue_rb_local_data_status` ON `contentCue` (`rb_local_data_status`);
CREATE INDEX `content_cue_rb_local_deleted` ON `contentCue` (`rb_local_deleted`);
CREATE INDEX `content_cue_rb_local_usn__i_d` ON `contentCue` (`rb_local_usn`, `ID`);
CREATE INDEX `content_file__content_i_d` ON `contentFile` (`ContentID`);
CREATE INDEX `content_file_rb_file_hash_dirty` ON `contentFile` (`rb_file_hash_dirty`);
CREATE INDEX `content_file_rb_file_size_dirty` ON `contentFile` (`rb_file_size_dirty`);
CREATE INDEX `content_file_rb_local_deleted_rb_in_progress_rb_local_file_status_rb_process_type_rb_priority` ON `contentFile` (
  `rb_local_deleted`, `rb_in_progress`, 
  `rb_local_file_status`, `rb_process_type`, 
  `rb_priority`
);
CREATE INDEX `content_file__u_u_i_d` ON `contentFile` (`UUID`);
CREATE INDEX `content_file_rb_data_status` ON `contentFile` (`rb_data_status`);
CREATE INDEX `content_file_rb_local_data_status` ON `contentFile` (`rb_local_data_status`);
CREATE INDEX `content_file_rb_local_deleted` ON `contentFile` (`rb_local_deleted`);
CREATE INDEX `content_file_rb_local_usn__i_d` ON `contentFile` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_active_censor__content_i_d` ON `djmdActiveCensor` (`ContentID`);
CREATE INDEX `djmd_active_censor__content_u_u_i_d` ON `djmdActiveCensor` (`ContentUUID`);
CREATE INDEX `djmd_active_censor__u_u_i_d` ON `djmdActiveCensor` (`UUID`);
CREATE INDEX `djmd_active_censor_rb_data_status` ON `djmdActiveCensor` (`rb_data_status`);
CREATE INDEX `djmd_active_censor_rb_local_data_status` ON `djmdActiveCensor` (`rb_local_data_status`);
CREATE INDEX `djmd_active_censor_rb_local_deleted` ON `djmdActiveCensor` (`rb_local_deleted`);
CREATE INDEX `djmd_active_censor_rb_local_usn__i_d` ON `djmdActiveCensor` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_album__name` ON `djmdAlbum` (`Name`);
CREATE INDEX `djmd_album__album_artist_i_d` ON `djmdAlbum` (`AlbumArtistID`);
CREATE INDEX `djmd_album__u_u_i_d` ON `djmdAlbum` (`UUID`);
CREATE INDEX `djmd_album_rb_data_status` ON `djmdAlbum` (`rb_data_status`);
CREATE INDEX `djmd_album_rb_local_data_status` ON `djmdAlbum` (`rb_local_data_status`);
CREATE INDEX `djmd_album_rb_local_deleted` ON `djmdAlbum` (`rb_local_deleted`);
CREATE INDEX `djmd_album_rb_local_usn__i_d` ON `djmdAlbum` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_artist__name` ON `djmdArtist` (`Name`);
CREATE INDEX `djmd_artist__u_u_i_d` ON `djmdArtist` (`UUID`);
CREATE INDEX `djmd_artist_rb_data_status` ON `djmdArtist` (`rb_data_status`);
CREATE INDEX `djmd_artist_rb_local_data_status` ON `djmdArtist` (`rb_local_data_status`);
CREATE INDEX `djmd_artist_rb_local_deleted` ON `djmdArtist` (`rb_local_deleted`);
CREATE INDEX `djmd_artist_rb_local_usn__i_d` ON `djmdArtist` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_category__u_u_i_d` ON `djmdCategory` (`UUID`);
CREATE INDEX `djmd_category_rb_data_status` ON `djmdCategory` (`rb_data_status`);
CREATE INDEX `djmd_category_rb_local_data_status` ON `djmdCategory` (`rb_local_data_status`);
CREATE INDEX `djmd_category_rb_local_deleted` ON `djmdCategory` (`rb_local_deleted`);
CREATE INDEX `djmd_category_rb_local_usn__i_d` ON `djmdCategory` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_cloud_property__u_u_i_d` ON `djmdCloudProperty` (`UUID`);
CREATE INDEX `djmd_cloud_property_rb_data_status` ON `djmdCloudProperty` (`rb_data_status`);
CREATE INDEX `djmd_cloud_property_rb_local_data_status` ON `djmdCloudProperty` (`rb_local_data_status`);
CREATE INDEX `djmd_cloud_property_rb_local_deleted` ON `djmdCloudProperty` (`rb_local_deleted`);
CREATE INDEX `djmd_cloud_property_rb_local_usn__i_d` ON `djmdCloudProperty` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_color__u_u_i_d` ON `djmdColor` (`UUID`);
CREATE INDEX `djmd_color_rb_data_status` ON `djmdColor` (`rb_data_status`);
CREATE INDEX `djmd_color_rb_local_data_status` ON `djmdColor` (`rb_local_data_status`);
CREATE INDEX `djmd_color_rb_local_deleted` ON `djmdColor` (`rb_local_deleted`);
CREATE INDEX `djmd_color_rb_local_usn__i_d` ON `djmdColor` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_content__master_d_b_i_d__master_song_i_d` ON `djmdContent` (`MasterDBID`, `MasterSongID`);
CREATE INDEX `djmd_content__genre_i_d` ON `djmdContent` (`GenreID`);
CREATE INDEX `djmd_content__label_i_d` ON `djmdContent` (`LabelID`);
CREATE INDEX `djmd_content__artist_i_d` ON `djmdContent` (`ArtistID`);
CREATE INDEX `djmd_content__remixer_i_d` ON `djmdContent` (`RemixerID`);
CREATE INDEX `djmd_content__org_artist_i_d` ON `djmdContent` (`OrgArtistID`);
CREATE INDEX `djmd_content__composer_i_d` ON `djmdContent` (`ComposerID`);
CREATE INDEX `djmd_content__album_i_d` ON `djmdContent` (`AlbumID`);
CREATE INDEX `djmd_content__key_i_d` ON `djmdContent` (`KeyID`);
CREATE INDEX `djmd_content_rb_local_deleted__service_i_d` ON `djmdContent` (`rb_local_deleted`, `ServiceID`);
CREATE INDEX `djmd_content_rb_local_deleted__file_type` ON `djmdContent` (`rb_local_deleted`, `FileType`);
CREATE INDEX `djmd_content_rb_local_deleted__bit_rate` ON `djmdContent` (`rb_local_deleted`, `BitRate`);
CREATE INDEX `djmd_content_rb_local_deleted__bit_depth` ON `djmdContent` (`rb_local_deleted`, `BitDepth`);
CREATE INDEX `djmd_content__u_u_i_d` ON `djmdContent` (`UUID`);
CREATE INDEX `djmd_content_rb_data_status` ON `djmdContent` (`rb_data_status`);
CREATE INDEX `djmd_content_rb_local_data_status` ON `djmdContent` (`rb_local_data_status`);
CREATE INDEX `djmd_content_rb_local_deleted` ON `djmdContent` (`rb_local_deleted`);
CREATE INDEX `djmd_content_rb_local_usn__i_d` ON `djmdContent` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_cue__content_i_d` ON `djmdCue` (`ContentID`);
CREATE INDEX `djmd_cue__content_u_u_i_d` ON `djmdCue` (`ContentUUID`);
CREATE INDEX `djmd_cue__u_u_i_d` ON `djmdCue` (`UUID`);
CREATE INDEX `djmd_cue_rb_data_status` ON `djmdCue` (`rb_data_status`);
CREATE INDEX `djmd_cue_rb_local_data_status` ON `djmdCue` (`rb_local_data_status`);
CREATE INDEX `djmd_cue_rb_local_deleted` ON `djmdCue` (`rb_local_deleted`);
CREATE INDEX `djmd_cue_rb_local_usn__i_d` ON `djmdCue` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_device__u_u_i_d` ON `djmdDevice` (`UUID`);
CREATE INDEX `djmd_device_rb_data_status` ON `djmdDevice` (`rb_data_status`);
CREATE INDEX `djmd_device_rb_local_data_status` ON `djmdDevice` (`rb_local_data_status`);
CREATE INDEX `djmd_device_rb_local_deleted` ON `djmdDevice` (`rb_local_deleted`);
CREATE INDEX `djmd_device_rb_local_usn__i_d` ON `djmdDevice` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_genre__name` ON `djmdGenre` (`Name`);
CREATE INDEX `djmd_genre__u_u_i_d` ON `djmdGenre` (`UUID`);
CREATE INDEX `djmd_genre_rb_data_status` ON `djmdGenre` (`rb_data_status`);
CREATE INDEX `djmd_genre_rb_local_data_status` ON `djmdGenre` (`rb_local_data_status`);
CREATE INDEX `djmd_genre_rb_local_deleted` ON `djmdGenre` (`rb_local_deleted`);
CREATE INDEX `djmd_genre_rb_local_usn__i_d` ON `djmdGenre` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_history__name` ON `djmdHistory` (`Name`);
CREATE INDEX `djmd_history__parent_i_d` ON `djmdHistory` (`ParentID`);
CREATE INDEX `djmd_history__u_u_i_d` ON `djmdHistory` (`UUID`);
CREATE INDEX `djmd_history_rb_data_status` ON `djmdHistory` (`rb_data_status`);
CREATE INDEX `djmd_history_rb_local_data_status` ON `djmdHistory` (`rb_local_data_status`);
CREATE INDEX `djmd_history_rb_local_deleted` ON `djmdHistory` (`rb_local_deleted`);
CREATE INDEX `djmd_history_rb_local_usn__i_d` ON `djmdHistory` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_hot_cue_banklist__name` ON `djmdHotCueBanklist` (`Name`);
CREATE INDEX `djmd_hot_cue_banklist__parent_i_d` ON `djmdHotCueBanklist` (`ParentID`);
CREATE INDEX `djmd_hot_cue_banklist__u_u_i_d` ON `djmdHotCueBanklist` (`UUID`);
CREATE INDEX `djmd_hot_cue_banklist_rb_data_status` ON `djmdHotCueBanklist` (`rb_data_status`);
CREATE INDEX `djmd_hot_cue_banklist_rb_local_data_status` ON `djmdHotCueBanklist` (`rb_local_data_status`);
CREATE INDEX `djmd_hot_cue_banklist_rb_local_deleted` ON `djmdHotCueBanklist` (`rb_local_deleted`);
CREATE INDEX `djmd_hot_cue_banklist_rb_local_usn__i_d` ON `djmdHotCueBanklist` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_key__scale_name` ON `djmdKey` (`ScaleName`);
CREATE INDEX `djmd_key__u_u_i_d` ON `djmdKey` (`UUID`);
CREATE INDEX `djmd_key_rb_data_status` ON `djmdKey` (`rb_data_status`);
CREATE INDEX `djmd_key_rb_local_data_status` ON `djmdKey` (`rb_local_data_status`);
CREATE INDEX `djmd_key_rb_local_deleted` ON `djmdKey` (`rb_local_deleted`);
CREATE INDEX `djmd_key_rb_local_usn__i_d` ON `djmdKey` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_label__name` ON `djmdLabel` (`Name`);
CREATE INDEX `djmd_label__u_u_i_d` ON `djmdLabel` (`UUID`);
CREATE INDEX `djmd_label_rb_data_status` ON `djmdLabel` (`rb_data_status`);
CREATE INDEX `djmd_label_rb_local_data_status` ON `djmdLabel` (`rb_local_data_status`);
CREATE INDEX `djmd_label_rb_local_deleted` ON `djmdLabel` (`rb_local_deleted`);
CREATE INDEX `djmd_label_rb_local_usn__i_d` ON `djmdLabel` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_menu_items__u_u_i_d` ON `djmdMenuItems` (`UUID`);
CREATE INDEX `djmd_menu_items_rb_data_status` ON `djmdMenuItems` (`rb_data_status`);
CREATE INDEX `djmd_menu_items_rb_local_data_status` ON `djmdMenuItems` (`rb_local_data_status`);
CREATE INDEX `djmd_menu_items_rb_local_deleted` ON `djmdMenuItems` (`rb_local_deleted`);
CREATE INDEX `djmd_menu_items_rb_local_usn__i_d` ON `djmdMenuItems` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_mixer_param__content_i_d` ON `djmdMixerParam` (`ContentID`);
CREATE INDEX `djmd_mixer_param__u_u_i_d` ON `djmdMixerParam` (`UUID`);
CREATE INDEX `djmd_mixer_param_rb_data_status` ON `djmdMixerParam` (`rb_data_status`);
CREATE INDEX `djmd_mixer_param_rb_local_data_status` ON `djmdMixerParam` (`rb_local_data_status`);
CREATE INDEX `djmd_mixer_param_rb_local_deleted` ON `djmdMixerParam` (`rb_local_deleted`);
CREATE INDEX `djmd_mixer_param_rb_local_usn__i_d` ON `djmdMixerParam` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_my_tag__parent_i_d` ON `djmdMyTag` (`ParentID`);
CREATE INDEX `djmd_my_tag__seq` ON `djmdMyTag` (`Seq`);
CREATE INDEX `djmd_my_tag__u_u_i_d` ON `djmdMyTag` (`UUID`);
CREATE INDEX `djmd_my_tag_rb_data_status` ON `djmdMyTag` (`rb_data_status`);
CREATE INDEX `djmd_my_tag_rb_local_data_status` ON `djmdMyTag` (`rb_local_data_status`);
CREATE INDEX `djmd_my_tag_rb_local_deleted` ON `djmdMyTag` (`rb_local_deleted`);
CREATE INDEX `djmd_my_tag_rb_local_usn__i_d` ON `djmdMyTag` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_playlist__parent_i_d` ON `djmdPlaylist` (`ParentID`);
CREATE INDEX `djmd_playlist__seq` ON `djmdPlaylist` (`Seq`);
CREATE INDEX `djmd_playlist__name` ON `djmdPlaylist` (`Name`);
CREATE INDEX `djmd_playlist__u_u_i_d` ON `djmdPlaylist` (`UUID`);
CREATE INDEX `djmd_playlist_rb_data_status` ON `djmdPlaylist` (`rb_data_status`);
CREATE INDEX `djmd_playlist_rb_local_data_status` ON `djmdPlaylist` (`rb_local_data_status`);
CREATE INDEX `djmd_playlist_rb_local_deleted` ON `djmdPlaylist` (`rb_local_deleted`);
CREATE INDEX `djmd_playlist_rb_local_usn__i_d` ON `djmdPlaylist` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_recommend_like__content_i_d1__content_i_d2` ON `djmdRecommendLike` (`ContentID1`, `ContentID2`);
CREATE INDEX `djmd_recommend_like__content_i_d2` ON `djmdRecommendLike` (`ContentID2`);
CREATE INDEX `djmd_recommend_like__u_u_i_d` ON `djmdRecommendLike` (`UUID`);
CREATE INDEX `djmd_recommend_like_rb_data_status` ON `djmdRecommendLike` (`rb_data_status`);
CREATE INDEX `djmd_recommend_like_rb_local_data_status` ON `djmdRecommendLike` (`rb_local_data_status`);
CREATE INDEX `djmd_recommend_like_rb_local_deleted` ON `djmdRecommendLike` (`rb_local_deleted`);
CREATE INDEX `djmd_recommend_like_rb_local_usn__i_d` ON `djmdRecommendLike` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_related_tracks__parent_i_d` ON `djmdRelatedTracks` (`ParentID`);
CREATE INDEX `djmd_related_tracks__seq` ON `djmdRelatedTracks` (`Seq`);
CREATE INDEX `djmd_related_tracks__name` ON `djmdRelatedTracks` (`Name`);
CREATE INDEX `djmd_related_tracks__u_u_i_d` ON `djmdRelatedTracks` (`UUID`);
CREATE INDEX `djmd_related_tracks_rb_data_status` ON `djmdRelatedTracks` (`rb_data_status`);
CREATE INDEX `djmd_related_tracks_rb_local_data_status` ON `djmdRelatedTracks` (`rb_local_data_status`);
CREATE INDEX `djmd_related_tracks_rb_local_deleted` ON `djmdRelatedTracks` (`rb_local_deleted`);
CREATE INDEX `djmd_related_tracks_rb_local_usn__i_d` ON `djmdRelatedTracks` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_sampler__parent_i_d` ON `djmdSampler` (`ParentID`);
CREATE INDEX `djmd_sampler__seq` ON `djmdSampler` (`Seq`);
CREATE INDEX `djmd_sampler__name` ON `djmdSampler` (`Name`);
CREATE INDEX `djmd_sampler__u_u_i_d` ON `djmdSampler` (`UUID`);
CREATE INDEX `djmd_sampler_rb_data_status` ON `djmdSampler` (`rb_data_status`);
CREATE INDEX `djmd_sampler_rb_local_data_status` ON `djmdSampler` (`rb_local_data_status`);
CREATE INDEX `djmd_sampler_rb_local_deleted` ON `djmdSampler` (`rb_local_deleted`);
CREATE INDEX `djmd_sampler_rb_local_usn__i_d` ON `djmdSampler` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_history__history_i_d` ON `djmdSongHistory` (`HistoryID`);
CREATE INDEX `djmd_song_history__content_i_d` ON `djmdSongHistory` (`ContentID`);
CREATE INDEX `djmd_song_history__u_u_i_d` ON `djmdSongHistory` (`UUID`);
CREATE INDEX `djmd_song_history_rb_data_status` ON `djmdSongHistory` (`rb_data_status`);
CREATE INDEX `djmd_song_history_rb_local_data_status` ON `djmdSongHistory` (`rb_local_data_status`);
CREATE INDEX `djmd_song_history_rb_local_deleted` ON `djmdSongHistory` (`rb_local_deleted`);
CREATE INDEX `djmd_song_history_rb_local_usn__i_d` ON `djmdSongHistory` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_hot_cue_banklist__hot_cue_banklist_i_d` ON `djmdSongHotCueBanklist` (`HotCueBanklistID`);
CREATE INDEX `djmd_song_hot_cue_banklist__content_i_d` ON `djmdSongHotCueBanklist` (`ContentID`);
CREATE INDEX `djmd_song_hot_cue_banklist__hot_cue_banklist_u_u_i_d` ON `djmdSongHotCueBanklist` (`HotCueBanklistUUID`);
CREATE INDEX `djmd_song_hot_cue_banklist__u_u_i_d` ON `djmdSongHotCueBanklist` (`UUID`);
CREATE INDEX `djmd_song_hot_cue_banklist_rb_data_status` ON `djmdSongHotCueBanklist` (`rb_data_status`);
CREATE INDEX `djmd_song_hot_cue_banklist_rb_local_data_status` ON `djmdSongHotCueBanklist` (`rb_local_data_status`);
CREATE INDEX `djmd_song_hot_cue_banklist_rb_local_deleted` ON `djmdSongHotCueBanklist` (`rb_local_deleted`);
CREATE INDEX `djmd_song_hot_cue_banklist_rb_local_usn__i_d` ON `djmdSongHotCueBanklist` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_my_tag__my_tag_i_d` ON `djmdSongMyTag` (`MyTagID`);
CREATE INDEX `djmd_song_my_tag__content_i_d` ON `djmdSongMyTag` (`ContentID`);
CREATE INDEX `djmd_song_my_tag__u_u_i_d` ON `djmdSongMyTag` (`UUID`);
CREATE INDEX `djmd_song_my_tag_rb_data_status` ON `djmdSongMyTag` (`rb_data_status`);
CREATE INDEX `djmd_song_my_tag_rb_local_data_status` ON `djmdSongMyTag` (`rb_local_data_status`);
CREATE INDEX `djmd_song_my_tag_rb_local_deleted` ON `djmdSongMyTag` (`rb_local_deleted`);
CREATE INDEX `djmd_song_my_tag_rb_local_usn__i_d` ON `djmdSongMyTag` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_playlist__playlist_i_d` ON `djmdSongPlaylist` (`PlaylistID`);
CREATE INDEX `djmd_song_playlist__content_i_d` ON `djmdSongPlaylist` (`ContentID`);
CREATE INDEX `djmd_song_playlist__u_u_i_d` ON `djmdSongPlaylist` (`UUID`);
CREATE INDEX `djmd_song_playlist_rb_data_status` ON `djmdSongPlaylist` (`rb_data_status`);
CREATE INDEX `djmd_song_playlist_rb_local_data_status` ON `djmdSongPlaylist` (`rb_local_data_status`);
CREATE INDEX `djmd_song_playlist_rb_local_deleted` ON `djmdSongPlaylist` (`rb_local_deleted`);
CREATE INDEX `djmd_song_playlist_rb_local_usn__i_d` ON `djmdSongPlaylist` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_related_tracks__related_tracks_i_d` ON `djmdSongRelatedTracks` (`RelatedTracksID`);
CREATE INDEX `djmd_song_related_tracks__content_i_d` ON `djmdSongRelatedTracks` (`ContentID`);
CREATE INDEX `djmd_song_related_tracks__u_u_i_d` ON `djmdSongRelatedTracks` (`UUID`);
CREATE INDEX `djmd_song_related_tracks_rb_data_status` ON `djmdSongRelatedTracks` (`rb_data_status`);
CREATE INDEX `djmd_song_related_tracks_rb_local_data_status` ON `djmdSongRelatedTracks` (`rb_local_data_status`);
CREATE INDEX `djmd_song_related_tracks_rb_local_deleted` ON `djmdSongRelatedTracks` (`rb_local_deleted`);
CREATE INDEX `djmd_song_related_tracks_rb_local_usn__i_d` ON `djmdSongRelatedTracks` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_sampler__sampler_i_d` ON `djmdSongSampler` (`SamplerID`);
CREATE INDEX `djmd_song_sampler__content_i_d` ON `djmdSongSampler` (`ContentID`);
CREATE INDEX `djmd_song_sampler__u_u_i_d` ON `djmdSongSampler` (`UUID`);
CREATE INDEX `djmd_song_sampler_rb_data_status` ON `djmdSongSampler` (`rb_data_status`);
CREATE INDEX `djmd_song_sampler_rb_local_data_status` ON `djmdSongSampler` (`rb_local_data_status`);
CREATE INDEX `djmd_song_sampler_rb_local_deleted` ON `djmdSongSampler` (`rb_local_deleted`);
CREATE INDEX `djmd_song_sampler_rb_local_usn__i_d` ON `djmdSongSampler` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_tag_list__content_i_d` ON `djmdSongTagList` (`ContentID`);
CREATE INDEX `djmd_song_tag_list__u_u_i_d` ON `djmdSongTagList` (`UUID`);
CREATE INDEX `djmd_song_tag_list_rb_data_status` ON `djmdSongTagList` (`rb_data_status`);
CREATE INDEX `djmd_song_tag_list_rb_local_data_status` ON `djmdSongTagList` (`rb_local_data_status`);
CREATE INDEX `djmd_song_tag_list_rb_local_deleted` ON `djmdSongTagList` (`rb_local_deleted`);
CREATE INDEX `djmd_song_tag_list_rb_local_usn__i_d` ON `djmdSongTagList` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_sort__u_u_i_d` ON `djmdSort` (`UUID`);
CREATE INDEX `djmd_sort_rb_data_status` ON `djmdSort` (`rb_data_status`);
CREATE INDEX `djmd_sort_rb_local_data_status` ON `djmdSort` (`rb_local_data_status`);
CREATE INDEX `djmd_sort_rb_local_deleted` ON `djmdSort` (`rb_local_deleted`);
CREATE INDEX `djmd_sort_rb_local_usn__i_d` ON `djmdSort` (`rb_local_usn`, `ID`);
CREATE INDEX `hot_cue_banklist_cue_rb_cue_count` ON `hotCueBanklistCue` (`rb_cue_count`);
CREATE INDEX `hot_cue_banklist_cue__hot_cue_banklist_i_d` ON `hotCueBanklistCue` (`HotCueBanklistID`);
CREATE INDEX `hot_cue_banklist_cue__u_u_i_d` ON `hotCueBanklistCue` (`UUID`);
CREATE INDEX `hot_cue_banklist_cue_rb_data_status` ON `hotCueBanklistCue` (`rb_data_status`);
CREATE INDEX `hot_cue_banklist_cue_rb_local_data_status` ON `hotCueBanklistCue` (`rb_local_data_status`);
CREATE INDEX `hot_cue_banklist_cue_rb_local_deleted` ON `hotCueBanklistCue` (`rb_local_deleted`);
CREATE INDEX `hot_cue_banklist_cue_rb_local_usn__i_d` ON `hotCueBanklistCue` (`rb_local_usn`, `ID`);
CREATE INDEX `setting_file_rb_file_hash_dirty` ON `settingFile` (`rb_file_hash_dirty`);
CREATE INDEX `setting_file_rb_file_size_dirty` ON `settingFile` (`rb_file_size_dirty`);
CREATE INDEX `setting_file__u_u_i_d` ON `settingFile` (`UUID`);
CREATE INDEX `setting_file_rb_data_status` ON `settingFile` (`rb_data_status`);
CREATE INDEX `setting_file_rb_local_data_status` ON `settingFile` (`rb_local_data_status`);
CREATE INDEX `setting_file_rb_local_deleted` ON `settingFile` (`rb_local_deleted`);
CREATE INDEX `setting_file_rb_local_usn__i_d` ON `settingFile` (`rb_local_usn`, `ID`);
CREATE INDEX `uuid_i_d_map__u_u_i_d` ON `uuidIDMap` (`UUID`);
CREATE INDEX `uuid_i_d_map_rb_data_status` ON `uuidIDMap` (`rb_data_status`);
CREATE INDEX `uuid_i_d_map_rb_local_data_status` ON `uuidIDMap` (`rb_local_data_status`);
CREATE INDEX `uuid_i_d_map_rb_local_deleted` ON `uuidIDMap` (`rb_local_deleted`);
CREATE INDEX `uuid_i_d_map_rb_local_usn__i_d` ON `uuidIDMap` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_playlist__attribute` ON `djmdPlaylist` (`Attribute`);
CREATE INDEX `djmd_cue__content_i_d_rb_local_deleted` ON `djmdCue` (`ContentID`, `rb_local_deleted`);
CREATE INDEX `djmd_mixer_param__content_i_d_rb_local_deleted` ON `djmdMixerParam` (`ContentID`, `rb_local_deleted`);
CREATE INDEX `djmd_song_history__content_i_d_rb_local_deleted` ON `djmdSongHistory` (`ContentID`, `rb_local_deleted`);
CREATE INDEX `djmd_song_my_tag__content_i_d_rb_local_deleted` ON `djmdSongMyTag` (`ContentID`, `rb_local_deleted`);
CREATE INDEX `djmd_song_playlist__content_i_d_rb_local_deleted` ON `djmdSongPlaylist` (`ContentID`, `rb_local_deleted`);
CREATE INDEX `image_file__table_name__target_u_u_i_d` ON `imageFile` (`TableName`, `TargetUUID`);
CREATE INDEX `image_file__table_name__target_i_d` ON `imageFile` (`TableName`, `TargetID`);
CREATE INDEX `image_file_rb_file_hash_dirty` ON `imageFile` (`rb_file_hash_dirty`);
CREATE INDEX `image_file_rb_file_size_dirty` ON `imageFile` (`rb_file_size_dirty`);
CREATE INDEX `image_file_rb_local_deleted_rb_in_progress_rb_local_file_status_rb_process_type_rb_priority` ON `imageFile` (
  `rb_local_deleted`, `rb_in_progress`, 
  `rb_local_file_status`, `rb_process_type`, 
  `rb_priority`
);
CREATE INDEX `image_file__u_u_i_d` ON `imageFile` (`UUID`);
CREATE INDEX `image_file_rb_data_status` ON `imageFile` (`rb_data_status`);
CREATE INDEX `image_file_rb_local_data_status` ON `imageFile` (`rb_local_data_status`);
CREATE INDEX `image_file_rb_local_deleted` ON `imageFile` (`rb_local_deleted`);
CREATE INDEX `image_file_rb_local_usn__i_d` ON `imageFile` (`rb_local_usn`, `ID`);
CREATE INDEX `djmd_song_my_tag__my_tag_i_d_rb_local_deleted__i_d` ON `djmdSongMyTag` (
  `MyTagID`, `rb_local_deleted`, `ID`
);

