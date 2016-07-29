package org.cloudfoundry.autoscaler.scheduler.util;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Calendar;
import java.util.Collections;
import java.util.List;
import java.util.TimeZone;

/**
 * Helper class for validating the data
 */
public class DataValidationHelper {

	/**
	 * Checks if the specified object is null.
	 * 
	 * @param object
	 * @return true if not null otherwise false
	 */
	public static boolean isNotNull(Object object) {
		if (object == null)
			return false;
		return true;
	}

	/**
	 * Checks if specified string is not empty (not null and not blank)
	 * 
	 * @param string
	 * @return true or false
	 */
	public static boolean isNotEmpty(String string) {
		if (isNotNull(string) && !string.isEmpty())
			return true;
		return false;
	}

	/**
	 * Checks if the timezone is valid
	 * 
	 * @param timeZoneId
	 * @return
	 */
	public static boolean isValidTimeZone(String timeZoneId) {
		if (isNotNull(timeZoneId)) {
			List<String> supportedTimeZones = Arrays.asList(DateHelper.supportedTimezones);
			if (supportedTimeZones.contains(timeZoneId)) {
				return true;
			}
			return false;
		} else {
			return false;
		}
	}

	/**
	 * Checks if the specified time in milli seconds is after now (current time).
	 * 
	 * @param timeInMillis
	 * @param timeZone
	 * 
	 * @return
	 */
	public static boolean isLaterThanNow(Long timeInMillis, TimeZone timeZone) {
		if (timeInMillis != null && timeZone != null) {
			Calendar calendar = Calendar.getInstance(timeZone);
			Long nowInMillis = calendar.getTimeInMillis();

			return timeInMillis > nowInMillis;
		}
		return false;
	}

	public static boolean isAfter(Long endTimeInMillis, Long startTimeInMillis) {
		if (isNotNull(startTimeInMillis) && isNotNull(endTimeInMillis)) {
			return endTimeInMillis > startTimeInMillis;
		} else {
			return false;
		}

	}

	/**
	 * This method is given a collection of SpecificDateScheduleDateTime (holding the schedule 
	 * identifier and its start date time and end date time). It traverses through the collection 
	 * to check if the the date time between different schedules overlap. If there is an overlap 
	 * then an error message is added to a collection and collection of messages is returned.
	 * 
	 * @param scheduleStartEndTimeList
	 * @return - List of date time overlap validation messages
	 */
	public static List<String[]> isNotOverlapForSpecificDate(
			List<SpecificDateScheduleDateTime> scheduleStartEndTimeList) {
		List<String[]> overlapDateTimeValidationErrorMsgList = new ArrayList<>();
		if (scheduleStartEndTimeList != null && !scheduleStartEndTimeList.isEmpty()) {

			Collections.sort(scheduleStartEndTimeList);

			for (int index = 0; index < scheduleStartEndTimeList.size() - 1; index++) {
				SpecificDateScheduleDateTime current = scheduleStartEndTimeList.get(index);
				SpecificDateScheduleDateTime next = scheduleStartEndTimeList.get(index + 1);

				// Check for date time overlaps and create a validation error message string array
				if (Long.compare(current.getStartDateTime(), next.getStartDateTime()) == 0) {

					// startDateTime values are equal, so an overlap. Set up a message for validation error
					String[] overlapDateTimeValidationErrorMsg = { ScheduleTypeEnum.SPECIFIC_DATE.getDescription(),
							current.getScheduleIdentifier(), "start_date start_time", next.getScheduleIdentifier(),
							"start_date start_time" };
					overlapDateTimeValidationErrorMsgList.add(overlapDateTimeValidationErrorMsg);
				} else if (Long.compare(current.getEndDateTime(), next.getStartDateTime()) >= 0) {// current startDateTime was earlier than next startDateTime

					// endDateTime of current is later than or equal to startDateTime of next. Set up a message for validation error
					String[] overlapDateTimeValidationErrorMsg = { ScheduleTypeEnum.SPECIFIC_DATE.getDescription(),
							current.getScheduleIdentifier(), "end_date end_time", next.getScheduleIdentifier(),
							"start_date start_time" };
					overlapDateTimeValidationErrorMsgList.add(overlapDateTimeValidationErrorMsg);
				}
			}
		}
		return overlapDateTimeValidationErrorMsgList;
	}

}