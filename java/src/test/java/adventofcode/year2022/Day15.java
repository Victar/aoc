package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

public class Day15 extends BaseTest {

	public static final int DAY = 15;
	//	public static  final int MAX = 20;
	public static final int MAX = 4000000;

	public static boolean canPresent(Point point, List<SensorBeacon> sbList, List<Point> beacons) {
		if (beacons.contains(point)) {
			return true;
		}
		for (SensorBeacon sb : sbList) {
			if (!sb.canPresent(point)) {
				return false;
			}
		}
		return true;
	}

	public static boolean canPresentGold(Point point, List<SensorBeacon> sbList, List<Point> beacons) {
		if (beacons.contains(point)) {
			return false;
		}
		for (SensorBeacon sb : sbList) {
			if (!sb.canPresent(point)) {
				return false;
			}
		}
		return true;
	}

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		List<SensorBeacon> sbList = new ArrayList<>();
		List<Point> beacons = new ArrayList<>();

		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			SensorBeacon sb = new SensorBeacon(input);
			sbList.add(sb);
			beacons.add(sb.getBeacon());
		}
		int y = 2000000;//10;
		//		int y = 10;

		int xStart = -10000000;
		int xEnd = 10000000;
		int count = 0;
		for (int x = xStart; x < xEnd; x++) {
			Point current = new Point(x, y);
			if (!canPresent(current, sbList, beacons)) {
				count++;
			}
		}
		System.out.println(count);

	}

	@Test public void runGold() throws Exception {
		List<SensorBeacon> sbList = new ArrayList<>();
		List<Point> beacons = new ArrayList<>();
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			SensorBeacon sb = new SensorBeacon(input);
			sbList.add(sb);
			beacons.add(sb.getBeacon());
		}
		for (SensorBeacon sensorBeacon : sbList) {
			Point current = sensorBeacon.checkOnBorder(sbList, beacons);
			if (current != null) {
				System.out.println(current);
				System.out.println(current.x * 4000000 + current.y);
				break;
			}
		}
	}

	@Data static class Point {

		long x;
		long y;

		public Point(long x, long y) {
			this.x = x;
			this.y = y;
		}

		public long geManhattan(Point p) {
			return Math.abs(this.x - p.x) + Math.abs(this.y - p.y);
		}
	}

	@Data class SensorBeacon {

		Point sensor;
		Point beacon;

		//		Sensor at x=8, y=7: closest beacon is at x=2, y=10
		public SensorBeacon(String input) {
			String[] spStr = input.split(": closest beacon is at x=");
			String[] sStr = spStr[0].substring(12).split(", y=");
			String[] bStr = spStr[1].split(", y=");
			this.sensor = new Point(Integer.parseInt(sStr[0]), Integer.parseInt(sStr[1]));
			this.beacon = new Point(Integer.parseInt(bStr[0]), Integer.parseInt(bStr[1]));
		}

		public boolean canPresent(Point point) {
			return sensor.geManhattan(point) > sensor.geManhattan(beacon);
		}

		public Point checkOnBorder(List<SensorBeacon> sbList, List<Point> beacons) {
			long radious = sensor.geManhattan(beacon);
			long startX = sensor.x - radious;
			long endX = sensor.x + radious;
			for (long x = startX; x <= endX; x++) {

				long yDiff = Math.abs(x - sensor.x);
				long y1 = sensor.y + (yDiff - radious);
				long y2 = sensor.y - (yDiff - radious);
				Point current1 = checkBorderPoint(sbList, beacons, new Point(x, y1));
				if (current1 != null) {
					return current1;
				}
				Point current2 = checkBorderPoint(sbList, beacons, new Point(x, y2));
				if (current2 != null) {
					return current2;
				}
			}
			return null;
		}

		public Point checkBorderPoint(List<SensorBeacon> sbList, List<Point> beacons, Point borderPoint) {
			for (long x = borderPoint.x - 1; x <= borderPoint.x + 1; x++) {
				for (long y = borderPoint.y - 1; y <= borderPoint.y + 1; y++) {
					if (x < MAX && y < MAX && x > 0 && y > 0) {
						Point current = new Point(x, y);
						if (Day15.canPresentGold(current, sbList, beacons)) {
							return current;
						}
					}
				}
			}
			return null;
		}

	}

}
